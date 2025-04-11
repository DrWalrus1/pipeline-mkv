package parsers

import (
	"servermakemkv/outputs"
	"strings"
)

const DISC_INFO_PREFIX = "CINFO:"

func ParseDiscInfo(input string) (*outputs.DiscInformation, error) {
	var discInfo outputs.DiscInformation

	trimmed, found := strings.CutPrefix(input, DISC_INFO_PREFIX)
	if !found {
		return errorPrefixNotFound[outputs.DiscInformation]()
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return errorNotEnoughValues[outputs.DiscInformation]()
	}

	discInfo.ID = split[0]
	discInfo.Code = split[1]
	discInfo.Value = split[2]

	return &discInfo, nil
}
