package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCurrentProgressTitleOutputSuccessfullyParse(t *testing.T) {
	expected := outputs.CurrentProgressTitleOutput{
		Code: "1",
		ID:   "1",
		Name: "Test",
	}

	input := "PRGC:1,1,Test"

	actual, err := parsers.ParseCurrentProgressTitleOutput(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestCurrentProgressTitleOutputFailsWhenMissingPrefix(t *testing.T) {
	input := "1,1,Test"

	actual, err := parsers.ParseCurrentProgressTitleOutput(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestCurrentProgressTitleOutputFailsWhenMissingValues(t *testing.T) {
	input := "1,1"

	actual, err := parsers.ParseCurrentProgressTitleOutput(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}
