package Requests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	stringBody := string(body)
	if strings.HasSuffix(stringBody, "null") {
		stringBody = stringBody[:len(stringBody)-4]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%d: %s", resp.StatusCode, stringBody))
	}

	return body, nil
}
