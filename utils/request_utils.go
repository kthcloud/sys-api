package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func IsGoodStatusCode(code int) bool {
	return int(code/100) == 2
}

func IsUserError(code int) bool {
	return int(code/100) == 4
}

func IsInternalError(code int) bool {
	return int(code/100) == 5
}

func setBearerTokenHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func setBasicAuthHeaders(req *http.Request, username string, password string) {
	req.SetBasicAuth(username, password)
}

func setJsonHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func doRequestInternal(req *http.Request) (*http.Response, error) {
	// do request
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to do http request. details: %s", err)
		return nil, err
	}

	// check if we received anything
	if res.Body == nil {
		err = fmt.Errorf("failed to open response. details: no body")
		return nil, err
	}

	return res, nil
}

func addParams(req *http.Request, params map[string]string) {
	values := req.URL.Query()
	for key, value := range params {
		values.Add(key, value)
	}
	req.URL.RawQuery = values.Encode()
}

func DoRequest(method string, url string, requestBody []byte, params map[string]string) (*http.Response, error) {
	// prepare request
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if params == nil {
		setJsonHeader(req)
	}
	addParams(req, params)
	return doRequestInternal(req)
}

func DoJsonGetRequest[T any](url string, params map[string]string) (*T, error) {
	res, err := DoRequest("GET", url, nil, params)
	if err != nil {
		return nil, err
	}

	body, err := ParseBody[T](res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DoRequestBearer(method string, url string, requestBody []byte, params map[string]string, token string) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if params == nil {
		setJsonHeader(req)
	}
	setBearerTokenHeaders(req, token)
	addParams(req, params)
	return doRequestInternal(req)
}

func DoRequestBasicAuth(method string, url string, requestBody []byte, params map[string]string, username string, password string) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if params == nil {
		setJsonHeader(req)
	}
	setBasicAuthHeaders(req, username, password)
	addParams(req, params)
	return doRequestInternal(req)
}

func ParseBody[T any](closer io.ReadCloser) (*T, error) {
	body, err := ReadBody(closer)
	if err != nil {
		return nil, err
	}
	defer CloseBody(closer)

	res, err := ParseJson[T](body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CloseBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Println("failed to close response body. details: ", err)
	}
}

func ReadBody(responseBody io.ReadCloser) ([]byte, error) {
	// read body
	body, err := io.ReadAll(responseBody)
	if err != nil {
		err = fmt.Errorf("failed to read response body. details: %s", err)
		return nil, err
	}
	return body, nil
}

func ParseJson[T any](data []byte) (*T, error) {
	var out T
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json data. details: %s", err)
	}
	return &out, nil
}
