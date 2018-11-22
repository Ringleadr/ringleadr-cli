package Requests

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetRequest(address string) ([]byte, error) {
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("500: " + string(body))
	}

	return body, nil
}
