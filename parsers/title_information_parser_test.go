package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestParseTitleInfoSuccessfully(t *testing.T) {
	expected := outputs.TitleInformation{
		ID:    "1",
		Code:  "CODE",
		Value: "Value",
	}

	input := "TINFO:1,CODE,Value"

	actual, err := parsers.ParseTitleInfo(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestParseTitleInfoFailsMissingPrefix(t *testing.T) {
	input := "1,CODE,Value"

	actual, err := parsers.ParseTitleInfo(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestParseTitleInfoFailsMissingValues(t *testing.T) {
	input := "TINFO:1,CODE"

	actual, err := parsers.ParseTitleInfo(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}
