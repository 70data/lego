package ucloud

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"k8s.io/klog/v2"
)

type Config struct {
	PublicKey       string `toml:"public_key"`
	PrivateKey      string `toml:"private_key"`
	BucketName      string `toml:"bucket_name"`
	FileHost        string `toml:"file_host"`
	BucketHost      string `toml:"bucket_host"`
	VerifyUploadMD5 bool   `toml:"verfiy_upload_md5"`
}

type Auth struct {
	publicKey  string
	privateKey string
}

func (A Auth) Authorization(method, bucket, key string, header http.Header) string {
	var sigData string
	method = strings.ToUpper(method)
	MD5 := header.Get("Content-MD5")
	contentType := header.Get("Content-Type")
	date := header.Get("Date")
	sigData = method + "\n" + MD5 + "\n" + contentType + "\n" + date + "\n"
	resource := "/" + bucket + "/" + key
	sigData += resource
	signature := A.signature(sigData)
	return "UCloud " + A.publicKey + ":" + signature
}

func (A Auth) AuthorizationPrivateURL(method, bucket, key, expires string, header http.Header) (string, string) {
	var sigData string
	method = strings.ToUpper(method)
	md5 := header.Get("Content-MD5")
	contentType := header.Get("Content-Type")
	sigData = method + "\n" + md5 + "\n" + contentType + "\n" + expires + "\n"
	resource := "/" + bucket + "/" + key
	sigData += resource
	signature := A.signature(sigData)
	return signature, A.publicKey
}

func (A Auth) signature(data string) string {
	mac := hmac.New(sha1.New, []byte(A.privateKey))
	mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func NewAuth(publicKey, privateKey string) Auth {
	return Auth{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

type UFileRequest struct {
	Auth               Auth
	BucketName         string
	Host               string
	Client             *http.Client
	Context            context.Context
	baseURL            *url.URL
	RequestHeader      http.Header
	LastResponseStatus int
	LastResponseHeader http.Header
	LastResponseBody   []byte
	verifyUploadMD5    bool
	lastResponse       *http.Response
}

func (u *UFileRequest) request(req *http.Request) error {
	resp, err := u.requestWithResp(req)
	if err != nil {
		klog.Infoln(err)
		return err
	}
	err = u.responseParse(resp)
	if err != nil {
		klog.Infoln(err)
		return err
	}
	if !VerifyHTTPCode(resp.StatusCode) {
		return err
	}

	return nil
}

func (u *UFileRequest) requestWithResp(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", "UFileGoSDK/2.02")
	resp, err = u.Client.Do(req.WithContext(u.Context))
	// If we got an error, and the context has been canceled, the context's error is probably more useful.
	if err != nil {
		select {
		case <-u.Context.Done():
			err = u.Context.Err()
		default:
		}
		return nil, err
	}
	reqHeader, _ := json.Marshal(req.Header)
	reqBody, _ := json.Marshal(req.Body)
	respHeader, _ := json.Marshal(resp.Header)
	respBody, _ := json.Marshal(resp.Body)
	klog.Infoln(req.URL, string(reqHeader), string(reqBody), resp.Status, string(respHeader), string(respBody))
	return resp, nil
}

func (u *UFileRequest) responseParse(resp *http.Response) error {
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		klog.Infoln(err)
		return err
	}
	defer resp.Body.Close()
	u.LastResponseStatus = resp.StatusCode
	u.LastResponseHeader = resp.Header
	u.LastResponseBody = resBody
	u.lastResponse = resp
	return nil
}

func (u *UFileRequest) genFileURL(keyName string) string {
	return u.baseURL.String() + keyName
}

// GetPrivateURL 获取私有空间的文件下载 URL
// keyName 表示传到 ufile 的文件名
// expiresDuation 表示下载链接的过期时间 从现在算起 24 * time.Hour 表示过期时间为一天
func (u *UFileRequest) GetPrivateURL(keyName string, expiresDuation time.Duration) string {
	t := time.Now()
	t = t.Add(expiresDuation)
	expires := strconv.FormatInt(t.Unix(), 10)
	signature, publicKey := u.Auth.AuthorizationPrivateURL("GET", u.BucketName, keyName, expires, http.Header{})
	query := url.Values{}
	query.Add("UCloudPublicKey", publicKey)
	query.Add("Signature", signature)
	query.Add("Expires", expires)
	reqURL := u.genFileURL(keyName)
	return reqURL + "?" + query.Encode()
}

func (u *UFileRequest) GetFilePrivateURL(keyName string) string {
	reqURL := u.GetPrivateURL(keyName, 365*24*time.Hour)
	return reqURL
}

func (u *UFileRequest) PutFile(filePath, keyName, mimeType string) error {
	file, err := openFile(filePath)
	if err != nil {
		klog.Infoln("File Path", filePath, err)
		return err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		klog.Infoln(err)
		return err
	}
	reqURL := u.genFileURL(keyName)
	if mimeType == "" {
		mimeType = getMimeType(file)
	}
	req, err := http.NewRequest("PUT", reqURL, bytes.NewBuffer(b))
	if err != nil {
		klog.Infoln(err)
		return err
	}
	req.Header.Add("Content-Type", mimeType)
	authorization := u.Auth.Authorization("PUT", u.BucketName, keyName, req.Header)
	req.Header.Add("authorization", authorization)
	fileSize := getFileSize(file)
	req.Header.Add("Content-Length", strconv.FormatInt(fileSize, 10))
	return u.request(req)
}

func (u *UFileRequest) DownloadFile(keyName string) error {
	reqURL := u.GetPrivateURL(keyName, 365*24*time.Hour)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return err
	}
	return u.request(req)
}

func NewFileRequest(config *Config, client *http.Client) *UFileRequest {
	config.BucketName = strings.TrimSpace(config.BucketName)
	config.FileHost = strings.TrimSpace(config.FileHost)
	req := newRequest(config.PublicKey, config.PrivateKey, config.BucketName, config.FileHost, client)
	req.verifyUploadMD5 = config.VerifyUploadMD5
	if req.baseURL.Scheme == "" {
		req.baseURL.Host = req.BucketName + "." + req.Host
		req.baseURL.Scheme = "http"
	}
	return req
}

func newRequest(publicKey, privateKey, bucket, host string, client *http.Client) *UFileRequest {
	req := new(UFileRequest)
	req.Auth = NewAuth(publicKey, privateKey)
	req.BucketName = bucket
	req.Host = strings.TrimSpace(host)
	req.baseURL = new(url.URL)
	req.baseURL.Host = req.Host
	// for default usage.
	req.baseURL.Path = "/"
	if client == nil {
		client = new(http.Client)
	}
	req.Client = client
	req.Context = context.TODO()
	return req
}

func VerifyHTTPCode(code int) bool {
	if code < http.StatusOK || code > http.StatusIMUsed {
		return false
	}
	return true
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func getMimeType(f *os.File) string {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	// revert file's seek
	defer func() { _, _ = f.Seek(0, 0) }()
	if err != nil {
		return "plain/text"
	}
	return http.DetectContentType(buffer)
}

func getFileSize(f *os.File) int64 {
	fi, err := f.Stat()
	if err != nil {
		panic(err.Error())
	}
	return fi.Size()
}

func LoadConfig(confPath string) (*Config, error) {
	c := new(Config)
	if _, err := toml.DecodeFile(confPath, &c); err != nil {
		klog.Infoln(err)
		return c, err
	}
	return c, nil
}
