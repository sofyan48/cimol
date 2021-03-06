package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Requester ...
type Requester struct{}

// RequesterHandler ...
func RequesterHandler() *Requester {
	return &Requester{}
}

// RequesterInterface ...
type RequesterInterface interface {
	GET(url, authorization string) ([]byte, error)
	POST(url, auth string, payload []byte) ([]byte, error)
	CLIENT(method, url string, payload []byte) (*http.Request, error)
}

// CLIENT ...
func (request *Requester) CLIENT(method, url string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return req, err
	}
	return req, nil
}

// GET request type get
func (request *Requester) GET(url, authorization string) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return body, nil
}

// POST request type post
func (request *Requester) POST(url, auth string, payload []byte) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+auth)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return body, nil
}
