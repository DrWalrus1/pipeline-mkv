package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestParseDiscInfoSuccessfully(t *testing.T) {
	expected := outputs.DiscInformation{
		ID:    "1",
		Code:  "CODE",
		Value: "Value",
	}

	input := "CINFO:1,CODE,Value"

	actual, err := parsers.ParseDiscInfo(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestParseDiscInfoFailsMissingPrefix(t *testing.T) {
	input := "1,CODE,Value"

	actual, err := parsers.ParseDiscInfo(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestParseDiscInfoFailsMissingValues(t *testing.T) {
	input := "CINFO:1,CODE"

	actual, err := parsers.ParseDiscInfo(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}
