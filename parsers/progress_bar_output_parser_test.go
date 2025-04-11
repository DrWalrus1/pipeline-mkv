package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestProgressBarOutputSuccessfulLine(t *testing.T) {
	expected := outputs.ProgressBarOutput{
		CurrentProgress: "1",
		TotalProgress:   "100",
		MaxProgress:     "200",
	}
	input := "PRGV:1,100,200"

	actual, err := parsers.ParseProgressBarOutput(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestProgressBarOutputMissingPrefix(t *testing.T) {
	input := "1,100,200"

	actual, err := parsers.ParseProgressBarOutput(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func TestProgressBarOutputIncomplete(t *testing.T) {
	input := "PRGV:1,100"

	actual, err := parsers.ParseProgressBarOutput(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}
