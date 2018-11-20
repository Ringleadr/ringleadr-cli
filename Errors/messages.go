package Errors

import (
	"errors"
)

func UnexpectedReponse() error {
	return errors.New("unexpected response received from host")
}

func FormatError() error {
	return errors.New("could not format response")
}
