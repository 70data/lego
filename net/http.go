package net

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

func HTTPGET(reqURL string) map[string]interface{} {
	req, _ := http.NewRequest("GET", reqURL, nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		klog.Infoln(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		klog.Infoln(berr)
	}
	responeDate := make(map[string]interface{})
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate
}

func HTTPPOST(reqURL, reqData string) map[string]interface{} {
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqData))
	req.Header.Add("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		klog.Infoln(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		klog.Infoln(berr)
		return nil
	}
	responeDate := make(map[string]interface{})
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate
}

func HTTPPUT(reqURL, reqData string) map[string]interface{} {
	req, _ := http.NewRequest("PUT", reqURL, strings.NewReader(reqData))
	req.Header.Add("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		klog.Infoln(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		klog.Infoln(berr)
		return nil
	}
	responeDate := make(map[string]interface{})
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate
}

func HTTPDELETE(reqURL, reqData string) map[string]interface{} {
	req, _ := http.NewRequest("DELETE", reqURL, nil)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		klog.Infoln(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		klog.Infoln(berr)
		return nil
	}
	responeDate := make(map[string]interface{})
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate
}
