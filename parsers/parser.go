package parsers

import (
	"errors"
	"servermakemkv/outputs"
	"strings"
)

type ParseError struct {
	Error error
}

type parserFunc func(string) (outputs.MakeMkvOutput, error)

var parsers = []struct {
	prefix string
	fn     func(string) (outputs.MakeMkvOutput, error)
}{
	{MESSAGE_OUTPUT_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseMessageOutput(s)
	}},
	{DRIVE_SCAN_MESSAGE_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseDriveScanMessage(s)
	}},
	{CURRENT_PROGRESS_TITLE_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseCurrentProgressTitleOutput(s)
	}},
	{DISC_INFO_OUTPUT_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseDiscInformationOutputMessage(s)
	}},
	{DISC_INFO_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseDiscInfo(s)
	}},
	{PROGRESS_BAR_OUTPUT_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseProgressBarOutput(s)
	}},
	{STREAM_INFO_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseStreamInfo(s)
	}},
	{TITLE_INFO_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseTitleInfo(s)
	}},
	{TOTAL_PROGRESS_TITLE_PREFIX, func(s string) (outputs.MakeMkvOutput, error) {
		return ParseTotalProgressTitleOutput(s)
	}},
}

func Parse(input string) (outputs.MakeMkvOutput, error) {
	for _, parser := range parsers {
		if strings.HasPrefix(input, parser.prefix) {
			return parser.fn(input)
		}
	}
	return nil, errors.New("Expected prefix was not found")
}
