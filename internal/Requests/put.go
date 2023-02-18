package Requests

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func PutRequest(address string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPut, address, body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%d: %s", resp.StatusCode, string(respBody)))
	}

	return respBody, nil
}
