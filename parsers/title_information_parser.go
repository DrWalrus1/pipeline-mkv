package parsers

import (
	"servermakemkv/outputs"
	"strings"
)

const TITLE_INFO_PREFIX = "TINFO:"

func ParseTitleInfo(input string) (*outputs.TitleInformation, error) {
	var titleInfo outputs.TitleInformation

	trimmed, found := strings.CutPrefix(input, TITLE_INFO_PREFIX)
	if !found {
		return errorPrefixNotFound[outputs.TitleInformation]()
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return errorNotEnoughValues[outputs.TitleInformation]()
	}

	titleInfo.ID = split[0]
	titleInfo.Code = split[1]
	titleInfo.Value = split[2]

	return &titleInfo, nil
}
