package parsers

import (
	"servermakemkv/outputs"
	"strconv"
	"strings"
)

const DRIVE_SCAN_MESSAGE_PREFIX = "DRV:"

func ParseDriveScanMessage(input string) (*outputs.DriveScanMessage, error) {
	var driveScanMessage outputs.DriveScanMessage

	trimmed, found := strings.CutPrefix(input, PROGRESS_BAR_OUTPUT_PREFIX)
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
