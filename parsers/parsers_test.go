package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMessageOutputSuccessfulLineNoParams(t *testing.T) {
	expected := outputs.MessageOutput{
		Code:           "1",
		Flags:          "test",
		ParameterCount: 0,
		RawMessage:     "Test Message",
		FormatMessage:  "Test Message",
	}

	input := "MSG:1,test,0,Test Message,Test Message"

	actual := parsers.ParseMessageOutput(input)

	assert.Equal(t, expected, actual)

}

func TestMessageOutputSuccessfulLineOneParam(t *testing.T) {
	expected := outputs.MessageOutput{
		Code:           "1",
		Flags:          "test",
		ParameterCount: 1,
		RawMessage:     "Test Message",
		FormatMessage:  "Test Message",
		MessageParams:  []string{"Test"},
	}

	input := "MSG:1,test,1,Test Message,Test Message,Test"

	actual := parsers.ParseMessageOutput(input)

	assert.Equal(t, expected, actual)

}

func TestMessageOutputSuccessfulLineThreeParams(t *testing.T) {
	expected := outputs.MessageOutput{
		Code:           "1",
		Flags:          "test",
		ParameterCount: 3,
		RawMessage:     "Test Message",
		FormatMessage:  "Test Message",
		MessageParams:  []string{"Test1", "Test2", "Test3"},
	}

	input := "MSG:1,test,3,Test Message,Test Message,Test1,Test2,Test3"

	actual := parsers.ParseMessageOutput(input)

	assert.Equal(t, expected, actual)

}

func TestMessageOutputFailsMissingPrefix(t *testing.T) {
	input := "1,test,3,Test Message,Test Message,Test1,Test2,Test3"

	actual := parsers.ParseMessageOutput(input)

	assert.Equal(t, nil, actual)
}

func TestMessageOutputParameterCountMismatchLessThanActual(t *testing.T) {
	expected := outputs.MessageOutput{
		Code:           "1",
		Flags:          "test",
		ParameterCount: 1,
		RawMessage:     "Test Message",
		FormatMessage:  "Test Message",
		MessageParams:  []string{"Test1"},
	}

	input := "MSG:1,test,1,Test Message,Test Message,Test1,Test2,Test3"

	actual := parsers.ParseMessageOutput(input)

	assert.Equal(t, expected, actual)
}

func TestMessageOutputParameterCountMismatchGreaterThanActualGrabsNoParams(t *testing.T) {
	expected := outputs.MessageOutput{
		Code:           "1",
		Flags:          "test",
		ParameterCount: 3,
		RawMessage:     "Test Message",
		FormatMessage:  "Test Message",
	}

	input := "MSG:1,test,3,Test Message,Test Message,Test1"

	actual := parsers.ParseMessageOutput(input)

	assert.Equal(t, expected, actual)
}

func TestProgressOutputSuccessfulLine(t *testing.T) {
	expected := outputs.ProgressBarOutput{
		CurrentProgress: "1",
		TotalProgress:   "100",
		MaxProgress:     "200",
	}
	input := "PRGV:1,100,200"

	actual := parsers.ParseProgressBarOutput(input)

	assert.Equal(t, expected, actual)
}

func TestProgressOutputMissingPrefix(t *testing.T) {
	input := "1,100,200"

	actual := parsers.ParseProgressBarOutput(input)

	assert.Equal(t, nil, actual)
}

func TestProgressOutputIncomplete(t *testing.T) {
	input := "PRGV:1,100"

	actual := parsers.ParseProgressBarOutput(input)

	assert.Equal(t, nil, actual)
}
