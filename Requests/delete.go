package Requests

import (
	"io/ioutil"
	"net/http"
)

func DeleteRequest(address string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, address, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
