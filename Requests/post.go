package Requests

import (
	"io"
	"io/ioutil"
	"net/http"
)

func PostRequest(address string, body io.Reader) ([]byte, error) {
	resp, err := http.Post(address, "application/json", body)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
