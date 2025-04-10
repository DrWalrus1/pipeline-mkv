package parsers

import (
	"servermakemkv/outputs"
	"strings"
)

const (
	progress_output_prefix = "PRGV:"
	delimiter              = ","
)

func ParseProgressString(input string) *outputs.ProgressOutput {
	var progressOutput outputs.ProgressOutput

	trimmed, found := strings.CutPrefix(input, progress_output_prefix)
	if !found {
		// TODO: figure out how to handle errors
		return nil
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil
	}
	progressOutput.CurrentProgress = split[0]
	progressOutput.TotalProgress = split[1]
	progressOutput.MaxProgress = split[2]
	return &progressOutput
}
