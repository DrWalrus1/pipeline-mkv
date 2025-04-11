package parsers

import (
	"servermakemkv/outputs"
	"strings"
)

const CURRENT_PROGRESS_TITLE_PREFIX = "PRGC:"

func ParseCurrentProgressTitleOutput(input string) (*outputs.CurrentProgressTitleOutput, error) {
	var currentProgressTitleOutput outputs.CurrentProgressTitleOutput

	trimmed, found := strings.CutPrefix(input, CURRENT_PROGRESS_TITLE_PREFIX)
	if !found {
		return errorPrefixNotFound[outputs.CurrentProgressTitleOutput]()
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return errorNotEnoughValues[outputs.CurrentProgressTitleOutput]()
	}

	currentProgressTitleOutput.Code = split[0]
	currentProgressTitleOutput.ID = split[1]
	currentProgressTitleOutput.Name = split[2]

	return &currentProgressTitleOutput, nil
}
