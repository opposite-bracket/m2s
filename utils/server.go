package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const HeaderTargetKey = "m2s-target"

type ReqInfo struct {
	Url        string              `json:"url"`
	Target     *url.URL            `json:"target"`
	Method     string              `json:"method"`
	Query      map[string][]string `json:"query,omitempty"`
	Body       interface{}         `json:"body,omitempty"`
	Header     map[string][]string `json:"header,omitempty"`
}

type ResInfo struct {
	Body   interface{}         `json:"body,omitempty"`
	Header map[string][]string `json:"header,omitempty"`
}

type E2EInfo struct {
	Req *ReqInfo `json:"req"`
	Res *ResInfo `json:"res"`
}

// LogReq sends e2e request info
// into log
func LogReq(e2eInfo E2EInfo) {

	id := time.Now().Format(time.RFC3339)
	fileName := fmt.Sprintf(
		"endpoint.%s.json",
		strings.Replace(id, ":", "-", 3),
	)

	LogOnError(
		"failed to log request: [error: %v]",
		WriteToJsonFile(fileName, e2eInfo),
	)
}

// ExtractBody converts body reader into interface
// (if any was found)
func ExtractBody(body io.ReadCloser) interface{} {

	var bodyMap interface{}
	err := json.NewDecoder(body).Decode(&bodyMap)

	if http.NoBody != body && err != nil {
		LogOnError(
			"failed to decode body: [error: %v]",
			err,
		)
	}

	return bodyMap
}

// ExtractReqInfo contains metadata about request made.
// This will log Url, Query, Method, Body & Header info
// into m2s log.
func ExtractReqInfo(req *http.Request) *ReqInfo {

	reqInfo := ReqInfo{
		Url:    req.URL.String(),
		Query:  req.URL.Query(),
		Method: req.Method,
		Body:   ExtractBody(req.Body),
		Header: req.Header,
	}

	if targetValue, err := url.ParseRequestURI(req.Header.Get(HeaderTargetKey)); err != nil {
		log.Printf("target url is invalid: [error: %s]", err)
	} else {
		reqInfo.Target = targetValue
		req.Header.Del(HeaderTargetKey)
	}

	return &reqInfo
}
