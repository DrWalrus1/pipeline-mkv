package parser

import (
	"errors"
	"servermakemkv/outputs"
	"strconv"
	"strings"
)

const (
	messageOutputPrefix        = "MSG:"
	driveScanMessagePrefix     = "DRV:"
	currentProgressTitlePrefix = "PRGC:"
	discInfoOutputPrefix       = "TCOUT:"
	discInfoPrefix             = "CINFO:"
	progressBarOutputPrefix    = "PRGV:"
	streamInfoPrefix           = "SINFO:"
	titleInfoPrefix            = "TINFO:"
	totalProgressTitlePrefix   = "PRGT:"
	delimiter                  = ","
)

var PrefixNotFound = errors.New("Prefix did not match expected")
var NotEnoughValues = errors.New("Not enough values found in input")
var EmptyInput = errors.New("input is empty")

type parserFunc func(string) (outputs.MakeMkvOutput, error)

var parsers = []struct {
	prefix string
	fn     func(string) (outputs.MakeMkvOutput, error)
}{
	{messageOutputPrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseMessageOutput(s)
	}},
	{driveScanMessagePrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseDriveScanMessage(s)
	}},
	{currentProgressTitlePrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseCurrentProgressTitleOutput(s)
	}},
	{discInfoOutputPrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseDiscInformationOutputMessage(s)
	}},
	{discInfoPrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseDiscInfo(s)
	}},
	{progressBarOutputPrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseProgressBarOutput(s)
	}},
	{streamInfoPrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseStreamInfo(s)
	}},
	{titleInfoPrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseTitleInfo(s)
	}},
	{totalProgressTitlePrefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseTotalProgressTitleOutput(s)
	}},
}

func Parse(input string) (outputs.MakeMkvOutput, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, EmptyInput
	}
	sanitised := strings.ReplaceAll(input, "\"", "")
	for _, parser := range parsers {
		if strings.HasPrefix(sanitised, parser.prefix) {
			return parser.fn(sanitised)
		}
	}
	return nil, PrefixNotFound
}

func parseMessageOutput(input string) (*outputs.MessageOutput, error) {
	var parsedMessage outputs.MessageOutput

	trimmed, _ := strings.CutPrefix(input, messageOutputPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 5 {
		return nil, NotEnoughValues
	}
	parsedMessage.Code = split[0]
	parsedMessage.Flags = split[1]
	parsedParamCount, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, NotEnoughValues
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

func parseDriveScanMessage(input string) (*outputs.DriveScanMessage, error) {
	var driveScanMessage outputs.DriveScanMessage

	trimmed, _ := strings.CutPrefix(input, driveScanMessagePrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 6 {
		return nil, NotEnoughValues
	}

	driveScanMessage.DriveIndex = split[0]
	driveScanMessage.Visible = split[1] == "1"
	driveScanMessage.Enabled = split[2] == "1"
	driveScanMessage.Flags = split[3]
	driveScanMessage.DriveName = split[4]
	driveScanMessage.DiscName = split[5]
	return &driveScanMessage, nil
}

func parseCurrentProgressTitleOutput(input string) (*outputs.CurrentProgressTitleOutput, error) {
	var currentProgressTitleOutput outputs.CurrentProgressTitleOutput

	trimmed, _ := strings.CutPrefix(input, currentProgressTitlePrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, NotEnoughValues
	}

	currentProgressTitleOutput.Code = split[0]
	currentProgressTitleOutput.ID = split[1]
	currentProgressTitleOutput.Name = split[2]

	return &currentProgressTitleOutput, nil
}

func parseDiscInformationOutputMessage(input string) (*outputs.DiscInformationOutputMessage, error) {
	var discInformationOutput outputs.DiscInformationOutputMessage

	trimmed, _ := strings.CutPrefix(input, discInfoOutputPrefix)

	titleCount, err := strconv.Atoi(trimmed)
	if err != nil {
		return nil, err
	}
	discInformationOutput.TitleCount = titleCount
	return &discInformationOutput, nil
}

func parseDiscInfo(input string) (*outputs.DiscInformation, error) {
	var discInfo outputs.DiscInformation

	trimmed, _ := strings.CutPrefix(input, discInfoPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, NotEnoughValues
	}

	discInfo.ID = split[0]
	discInfo.Code = split[1]
	discInfo.Value = split[2]

	return &discInfo, nil
}

func parseProgressBarOutput(input string) (*outputs.ProgressBarOutput, error) {
	var progressOutput outputs.ProgressBarOutput

	trimmed, _ := strings.CutPrefix(input, progressBarOutputPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, NotEnoughValues
	}
	progressOutput.CurrentProgress = split[0]
	progressOutput.TotalProgress = split[1]
	progressOutput.MaxProgress = split[2]
	return &progressOutput, nil
}

func parseStreamInfo(input string) (*outputs.StreamInformation, error) {
	var streamInfo outputs.StreamInformation

	trimmed, found := strings.CutPrefix(input, streamInfoPrefix)
	if !found {
		return nil, PrefixNotFound
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, NotEnoughValues
	}

	streamInfo.ID = split[0]
	streamInfo.Code = split[1]
	streamInfo.Value = split[2]

	return &streamInfo, nil
}

func parseTitleInfo(input string) (*outputs.TitleInformation, error) {
	var titleInfo outputs.TitleInformation

	trimmed, _ := strings.CutPrefix(input, titleInfoPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, NotEnoughValues
	}

	titleInfo.ID = split[0]
	titleInfo.Code = split[1]
	titleInfo.Value = split[2]

	return &titleInfo, nil
}

func parseTotalProgressTitleOutput(input string) (*outputs.TotalProgressTitleOutput, error) {
	var currentProgressTitleOutput outputs.TotalProgressTitleOutput

	trimmed, _ := strings.CutPrefix(input, totalProgressTitlePrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, NotEnoughValues
	}

	currentProgressTitleOutput.Code = split[0]
	currentProgressTitleOutput.ID = split[1]
	currentProgressTitleOutput.Name = split[2]

	return &currentProgressTitleOutput, nil

}
