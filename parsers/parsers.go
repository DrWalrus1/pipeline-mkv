package parsers

import (
	"servermakemkv/outputs"
	"strconv"
	"strings"
)

const (
	MESSAGE_OUTPUT_PREFIX         = "MSG:"
	CURRENT_PROGRESS_TITLE_PREFIX = "PRGC:"
	TOTAL_PROGRESS_TITLE_PREFIX   = "PRGT:"
	PROGRESS_BAR_OUTPUT_PREFIX    = "PRGV:"
	DRIVE_SCAN_MESSAGE_PREFIX     = "DRV:"
	DISC_INFO_OUTPUT_PREFIX       = "TCOUT:"
	DISC_INFO_PREFIX              = "CINFO:"
	TITLE_INFO_PREFIX             = "TINFO:"
	STREAM_INFO_PREFIX            = "SINFO:"
	delimiter                     = ","
)

func ParseMessageOutput(input string) *outputs.MessageOutput {
	var parsedMessage outputs.MessageOutput

	trimmed, found := strings.CutPrefix(input, MESSAGE_OUTPUT_PREFIX)
	if !found {
		return nil
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 5 {
		return nil
	}
	parsedMessage.Code = split[0]
	parsedMessage.Flags = split[1]
	parsedParamCount, err := strconv.Atoi(split[2])
	if err != nil {
		return nil
	}
	parsedMessage.ParameterCount = parsedParamCount
	parsedMessage.RawMessage = split[3]
	parsedMessage.FormatMessage = split[4]
	const splitOffset = 5
	ifThereAreParamsAfterOffset := func() bool {
		return len(split) > splitOffset
	}
	doParamsExist := func() bool {
		return parsedMessage.ParameterCount > 0
	}
	paramsDoNotExceedSplitBounds := func() bool {
		return parsedMessage.ParameterCount+splitOffset-1 < len(split)
	}
	if ifThereAreParamsAfterOffset() && doParamsExist() && paramsDoNotExceedSplitBounds() {
		parsedMessage.MessageParams = make([]string, parsedMessage.ParameterCount)
		for i := range parsedMessage.ParameterCount {
			parsedMessage.MessageParams[i] = split[i+splitOffset]
		}
	}

	return &parsedMessage
}

func ParseProgressBarOutput(input string) *outputs.ProgressBarOutput {
	var progressOutput outputs.ProgressBarOutput

	trimmed, found := strings.CutPrefix(input, PROGRESS_BAR_OUTPUT_PREFIX)
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
