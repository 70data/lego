package utils

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// HTTPGetReturnByte compatible http & https return byte with custom header
func HTTPGetReturnByte(reqURL string, header map[string]string) []byte {
	req, _ := http.NewRequest("GET", reqURL, nil)
	for key, value := range header {
		req.Header.Set(key, value)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
		Timeout:   300 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if berr != nil {
		log.Println(berr)
	}
	return resBody
}

// HTTPGet compatible http & https
func HTTPGet(reqURL string) map[string]interface{} {
	req, _ := http.NewRequest("GET", reqURL, nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
		Timeout:   300 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if berr != nil {
		log.Println(berr)
	}
	responeDate := make(map[string]interface{})
	json.Unmarshal(resBody, &responeDate)
	return responeDate
}

// HTTPPost is post func
func HTTPPost(reqURL, reqData string) map[string]interface{} {
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqData))
	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 300 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if berr != nil {
		log.Println(berr)
	}
	responeDate := make(map[string]interface{})
	json.Unmarshal(resBody, &responeDate)
	return responeDate
}

// HTTPSaltPost is post func with salt
func HTTPSaltPost(reqURL, reqData string) map[string]interface{} {
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqData))
	req.Header.Set("Content-Type", "application/json")
	// make salt
	timeNow := TimeUnix()
	req.Header.Set("NXOS-ts", timeNow)
	bodyMD := MakeMD(reqData)
	tokenNaive := bodyMD + timeNow
	bodyToken := MakeMD(tokenNaive)
	req.Header.Set("NXOS-token", bodyToken)
	c := &http.Client{
		Timeout: 300 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if berr != nil {
		log.Println(berr)
	}
	responeDate := make(map[string]interface{})
	json.Unmarshal(resBody, &responeDate)
	return responeDate
}

// HTTPPut is post func
func HTTPPut(reqURL, reqData string) map[string]interface{} {
	req, _ := http.NewRequest("PUT", reqURL, strings.NewReader(reqData))
	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 300 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if berr != nil {
		log.Println(berr)
	}
	responeDate := make(map[string]interface{})
	json.Unmarshal(resBody, &responeDate)
	return responeDate
}

// HTTPDelete is delete func
func HTTPDelete(reqURL, reqData string) map[string]interface{} {
	req, _ := http.NewRequest("DELETE", reqURL, nil)
	c := &http.Client{
		Timeout: 300 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if berr != nil {
		log.Println(berr)
	}
	responeDate := make(map[string]interface{})
	json.Unmarshal(resBody, &responeDate)
	return responeDate
}
