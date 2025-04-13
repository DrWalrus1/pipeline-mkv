package parser

import (
	"errors"
	"servermakemkv/outputs"
)

func errorPrefixNotFound[T outputs.MakeMkvOutput]() (*T, error) {
	return nil, errors.New("Prefix did not match expected")
}

func errorNotEnoughValues[T outputs.MakeMkvOutput]() (*T, error) {
	return nil, errors.New("Not enough values found in input")
}
