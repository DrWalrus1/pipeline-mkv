package parsers

import (
	"servermakemkv/outputs"
	"strconv"
	"strings"
)

const MESSAGE_OUTPUT_PREFIX = "MSG:"

func ParseMessageOutput(input string) (*outputs.MessageOutput, error) {
	var parsedMessage outputs.MessageOutput

	trimmed, found := strings.CutPrefix(input, MESSAGE_OUTPUT_PREFIX)
	if !found {
		return errorPrefixNotFound[outputs.MessageOutput]()
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 5 {
		return errorNotEnoughValues[outputs.MessageOutput]()
	}
	parsedMessage.Code = split[0]
	parsedMessage.Flags = split[1]
	parsedParamCount, err := strconv.Atoi(split[2])
	if err != nil {
		return errorNotEnoughValues[outputs.MessageOutput]()
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

	return &parsedMessage, nil
}
