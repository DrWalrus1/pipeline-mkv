package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestProgressOutputSuccessfulLine(t *testing.T) {
	expected := outputs.ProgressOutput{
		CurrentProgress: "1",
		TotalProgress:   "100",
		MaxProgress:     "200",
	}
	input := "PRGV:1,100,200"

	actual := parsers.ParseProgressString(input)

	assert.Equal(t, expected, actual)
}

func TestProgressOutputMissingPrefix(t *testing.T) {
	input := "1,100,200"

	actual := parsers.ParseProgressString(input)

	assert.Equal(t, nil, actual)
}

func TestProgressOutputIncomplete(t *testing.T) {
	input := "PRGV:1,100"

	actual := parsers.ParseProgressString(input)

	assert.Equal(t, nil, actual)
}
