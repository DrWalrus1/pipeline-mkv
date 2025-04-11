package parsers

import (
	"servermakemkv/outputs"
	"strings"
)

const STREAM_INFO_PREFIX = "SINFO:"

func ParseStreamInfo(input string) (*outputs.StreamInformation, error) {
	var streamInfo outputs.StreamInformation

	trimmed, found := strings.CutPrefix(input, STREAM_INFO_PREFIX)
	if !found {
		return errorPrefixNotFound[outputs.StreamInformation]()
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return errorNotEnoughValues[outputs.StreamInformation]()
	}

	streamInfo.ID = split[0]
	streamInfo.Code = split[1]
	streamInfo.Value = split[2]

	return &streamInfo, nil
}
