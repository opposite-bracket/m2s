package utils

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type m2sClient struct {
	*http.Client
}

func (c *m2sClient) MakeRequest(reqInfo *ReqInfo) (*ResInfo, error) {

	req, err := http.NewRequest(reqInfo.Method, reqInfo.Target.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	var response interface{}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Failed to read content: %v", err)
		return nil, err
	}
	resInfo := ResInfo{
		Header: res.Header,
		Body: response,
	}

	return &resInfo, nil
}

func NewClient() *m2sClient {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: config,
	}

	return &m2sClient{
		Client: &http.Client{Transport: transport},
	}
}
