package parsers

import "errors"

func errorPrefixNotFound[T interface{}]() (*T, error) {
	return nil, errors.New("Prefix did not match expected")
}

func errorNotEnoughValues[T interface{}]() (*T, error) {
	return nil, errors.New("Not enough values found in input")
}
