package Requests

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

func PostRequest(address string, body io.Reader) ([]byte, error) {
	resp, err := http.Post(address, "application/json", body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("500: " + string(respBody))
	}

	return respBody, nil
}
