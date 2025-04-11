package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestParseStreamInfoSuccessfully(t *testing.T) {
	expected := outputs.StreamInformation{
		ID:    "1",
		Code:  "CODE",
		Value: "Value",
	}

	input := "SINFO:1,CODE,Value"

	actual, err := parsers.ParseStreamInfo(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestParseStreamInfoFailsMissingPrefix(t *testing.T) {
	input := "1,CODE,Value"

	actual, err := parsers.ParseStreamInfo(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestParseStreamInfoFailsMissingValues(t *testing.T) {
	input := "SINFO:1,CODE"

	actual, err := parsers.ParseStreamInfo(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}
