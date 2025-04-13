package parser

import (
	"errors"
	"servermakemkv/outputs"
	"strconv"
	"strings"
)

const delimiter = ","

type ParseError struct {
	Error error
}

type parserFunc func(string) (outputs.MakeMkvOutput, error)

var parsers = []struct {
	prefix string
	fn     func(string) (outputs.MakeMkvOutput, error)
}{
	{message_output_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseMessageOutput(s)
	}},
	{drive_scan_message_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseDriveScanMessage(s)
	}},
	{current_progress_title_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseCurrentProgressTitleOutput(s)
	}},
	{disc_info_output_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseDiscInformationOutputMessage(s)
	}},
	{disc_info_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseDiscInfo(s)
	}},
	{progress_bar_output_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseProgressBarOutput(s)
	}},
	{stream_info_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseStreamInfo(s)
	}},
	{title_info_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseTitleInfo(s)
	}},
	{total_progress_title_prefix, func(s string) (outputs.MakeMkvOutput, error) {
		return parseTotalProgressTitleOutput(s)
	}},
}

func Parse(input string) (outputs.MakeMkvOutput, error) {
	for _, parser := range parsers {
		if strings.HasPrefix(input, parser.prefix) {
			return parser.fn(input)
		}
	}
	return nil, errors.New("Expected prefix was not found")
}

const message_output_prefix = "MSG:"

func parseMessageOutput(input string) (*outputs.MessageOutput, error) {
	var parsedMessage outputs.MessageOutput

	trimmed, found := strings.CutPrefix(input, message_output_prefix)
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

const drive_scan_message_prefix = "DRV:"

func parseDriveScanMessage(input string) (*outputs.DriveScanMessage, error) {
	var driveScanMessage outputs.DriveScanMessage

	trimmed, found := strings.CutPrefix(input, progress_bar_output_prefix)
	if !found {
		return errorPrefixNotFound[outputs.DriveScanMessage]()
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 6 {
		return errorNotEnoughValues[outputs.DriveScanMessage]()
	}

	driveScanMessage.DriveIndex = split[0]
	visible, err := strconv.ParseBool(split[1])
	if err != nil {
		return nil, err
	}
	driveScanMessage.Visible = visible
	enabled, err := strconv.ParseBool(split[2])
	if err != nil {
		return nil, err
	}
	driveScanMessage.Enabled = enabled
	driveScanMessage.Flags = split[3]
	driveScanMessage.DriveName = split[4]
	driveScanMessage.DiscName = split[5]
	return &driveScanMessage, nil
}

const current_progress_title_prefix = "PRGC:"

func parseCurrentProgressTitleOutput(input string) (*outputs.CurrentProgressTitleOutput, error) {
	var currentProgressTitleOutput outputs.CurrentProgressTitleOutput

	trimmed, found := strings.CutPrefix(input, current_progress_title_prefix)
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

const disc_info_output_prefix = "TCOUT:"

func parseDiscInformationOutputMessage(input string) (*outputs.DiscInformationOutputMessage, error) {
	var discInformationOutput outputs.DiscInformationOutputMessage

	trimmed, found := strings.CutPrefix(input, disc_info_output_prefix)
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

const disc_info_prefix = "CINFO:"

func parseDiscInfo(input string) (*outputs.DiscInformation, error) {
	var discInfo outputs.DiscInformation

	trimmed, found := strings.CutPrefix(input, disc_info_prefix)
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

const progress_bar_output_prefix = "PRGV:"

func parseProgressBarOutput(input string) (*outputs.ProgressBarOutput, error) {
	var progressOutput outputs.ProgressBarOutput

	trimmed, found := strings.CutPrefix(input, progress_bar_output_prefix)
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

const stream_info_prefix = "SINFO:"

func parseStreamInfo(input string) (*outputs.StreamInformation, error) {
	var streamInfo outputs.StreamInformation

	trimmed, found := strings.CutPrefix(input, stream_info_prefix)
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

const title_info_prefix = "TINFO:"

func parseTitleInfo(input string) (*outputs.TitleInformation, error) {
	var titleInfo outputs.TitleInformation

	trimmed, found := strings.CutPrefix(input, title_info_prefix)
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

const total_progress_title_prefix = "PRGT:"

func parseTotalProgressTitleOutput(input string) (*outputs.TotalProgressTitleOutput, error) {
	var currentProgressTitleOutput outputs.TotalProgressTitleOutput

	trimmed, found := strings.CutPrefix(input, total_progress_title_prefix)
	if !found {
		return nil, errors.New("Prefix did not match expected")
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, errors.New("Not enough values found")
	}

	currentProgressTitleOutput.Code = split[0]
	currentProgressTitleOutput.ID = split[1]
	currentProgressTitleOutput.Name = split[2]

	return &currentProgressTitleOutput, nil

}
