package parsers

import (
	"errors"
	"servermakemkv/outputs"
	"strings"
)

const TOTAL_PROGRESS_TITLE_PREFIX = "PRGT:"

func ParseTotalProgressTitleOutput(input string) (*outputs.TotalProgressTitleOutput, error) {
	var currentProgressTitleOutput outputs.TotalProgressTitleOutput

	trimmed, found := strings.CutPrefix(input, TOTAL_PROGRESS_TITLE_PREFIX)
	if !found {
		// TODO: figure out how to handle errors
		return nil, errors.New("Prefix did not match expected")
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, errors.New("Not enough values found")
	}

	currentProgressTitleOutput.Code = split[0]
	currentProgressTitleOutput.ID = split[1]
	currentProgressTitleOutput.Name = split[2]

	return &currentProgressTitleOutput, nil

}
