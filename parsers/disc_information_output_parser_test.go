package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestParseDiscInformationOutputSucessfully(t *testing.T) {
	expected := outputs.DiscInformationOutputMessage{
		TitleCount: 1,
	}

	input := "TCOUT:1"

	actual, err := parsers.ParseDiscInformationOutputMessage(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestParsedDiscInformationOutputFailsMissingPrefix(t *testing.T) {
	input := "1"

	actual, err := parsers.ParseDiscInformationOutputMessage(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestParsedDiscInformationOutputFailsMissingValues(t *testing.T) {
	input := "TCOUT:"
	actual, err := parsers.ParseDiscInformationOutputMessage(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestParsedDiscInformationOutputFailsNotValidValue(t *testing.T) {
	input := "TCOUT:THISISNOTANUMBER"

	actual, err := parsers.ParseDiscInformationOutputMessage(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)

}
