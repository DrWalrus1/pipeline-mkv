package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTotalTitleOutputSuccessfullyParse(t *testing.T) {
	expected := outputs.TotalProgressTitleOutput{
		Code: "1",
		ID:   "1",
		Name: "Test",
	}

	input := "PRGT:1,1,Test"

	actual, err := parsers.ParseTotalProgressTitleOutput(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestTotalTitleOutputFailsWhenMissingPrefix(t *testing.T) {
	input := "1,1,Test"

	actual, err := parsers.ParseTotalProgressTitleOutput(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestTotalTitleOutputFailsWhenMissingValues(t *testing.T) {
	input := "1,1"

	actual, err := parsers.ParseTotalProgressTitleOutput(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}
