package parsers

import (
	"servermakemkv/outputs"
	"strconv"
	"strings"
)

const DISC_INFO_OUTPUT_PREFIX = "TCOUT:"

func ParseDiscInformationOutputMessage(input string) (*outputs.DiscInformationOutputMessage, error) {
	var discInformationOutput outputs.DiscInformationOutputMessage

	trimmed, found := strings.CutPrefix(input, DISC_INFO_OUTPUT_PREFIX)
	if !found {
		return errorPrefixNotFound[outputs.DiscInformationOutputMessage]()
	}

	titleCount, err := strconv.Atoi(trimmed)
	if err != nil {
		return nil, err
	}
	discInformationOutput.TitleCount = titleCount
	return &discInformationOutput, nil
}
