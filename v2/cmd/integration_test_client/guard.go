package main

import "errors"

func NotEmptyArray(argsLen int) error {
	if argsLen == 0 {
		return errors.New("empty list received from nuclei-server")
	}
	return nil
}

func NotEmpty(result string) error {
	if len(result) == 0 {
		return errors.New("empty data received from nuclei-server")
	}
	return nil
}
