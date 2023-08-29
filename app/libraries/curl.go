package libraries

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type (
	Curl struct {
		Url         string
		Method      string
		FormData    map[string]string
		JsonData    string
		BodyRaw     string
		Headers     map[string]string
		BearerToken string
		BasicAuth   BasicAuth
		//Options  map[string]string
	}

	BasicAuth struct {
		Username string
		Password string
	}
)

func (cr *Curl) Execute() ([]byte, bool) {
	var bodyResp []byte
	var payload string = ""
	var contentType string = ""

	// initiate http client
	client := &http.Client{}

	// add body payload
	if cr.FormData != nil { // form-data type
		item := 0
		for key, value := range cr.FormData {
			field := key + "=" + value
			if item > 0 {
				field = "&" + field
			}
			payload = payload + field
			item++
		}
		contentType = "application/x-www-form-urlencoded"
	} else if cr.JsonData != "" { // json payload type
		payload = cr.JsonData
		contentType = "application/json"
	} else { // raw content type
		payload = cr.BodyRaw
	}

	// build request
	req, err := http.NewRequest(
		cr.Method,
		cr.Url,
		strings.NewReader(payload),
	)

	if err != nil {
		fmt.Println(err)
		return bodyResp, false
	}

	// set content type
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	// set bearer token
	if cr.BearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+cr.BearerToken)
	}

	// set custom header
	for key, value := range cr.Headers {
		//fmt.Println(key + " : " + value)
		req.Header.Add(key, value)
	}
	resApi, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return bodyResp, false
	}
	defer resApi.Body.Close()

	body, err := io.ReadAll(resApi.Body)

	if err != nil {
		fmt.Println(err)
		return bodyResp, false
	}

	bodyResp = body
	return bodyResp, true
}
