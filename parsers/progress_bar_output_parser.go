package parsers

import (
	"servermakemkv/outputs"
	"strings"
)

const PROGRESS_BAR_OUTPUT_PREFIX = "PRGV:"

func ParseProgressBarOutput(input string) (*outputs.ProgressBarOutput, error) {
	var progressOutput outputs.ProgressBarOutput

	trimmed, found := strings.CutPrefix(input, PROGRESS_BAR_OUTPUT_PREFIX)
	if !found {
		return errorPrefixNotFound[outputs.ProgressBarOutput]()
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return errorNotEnoughValues[outputs.ProgressBarOutput]()
	}
	progressOutput.CurrentProgress = split[0]
	progressOutput.TotalProgress = split[1]
	progressOutput.MaxProgress = split[2]
	return &progressOutput, nil
}
