package parsers_test

import (
	"servermakemkv/outputs"
	"servermakemkv/parsers"
	"testing"

	"github.com/go-playground/assert/v2"
)

func DriveScanMessageParsedSuccessfully(t *testing.T) {
	expected := outputs.DriveScanMessage{
		DriveIndex: "1",
		Visible:    true,
		Enabled:    true,
		Flags:      "Flags",
		DriveName:  "Drive1",
		DiscName:   "Disc1",
	}

	input := "DRV:1,true,true,Flags,Drive1,Disc1"

	actual, err := parsers.ParseDriveScanMessage(input)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func DriveScanMessageParseFailMissingPrefix(t *testing.T) {
	input := "1,true,true,Flags,Drive1,Disc1"

	actual, err := parsers.ParseDriveScanMessage(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func DriveScanMessageParseFailNotEnoughValues(t *testing.T) {
	input := "DRV:1,true,true,Flags,Drive1"

	actual, err := parsers.ParseDriveScanMessage(input)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, actual)
}

func DriveScanMessageFailNotValidBooleans(t *testing.T) {
	input1 := "DRV:1,NOTABOOL,true,Flags,Drive1,Disc1"
	input2 := "DRV:1,true,NOTABOOL,Flags,Drive1,Disc1"
	input3 := "DRV:1,NOTABOOL,NOTABOOL,Flags,Drive1,Disc1"

	actual1, err1 := parsers.ParseDriveScanMessage(input1)

	assert.NotEqual(t, nil, err1)
	assert.Equal(t, nil, actual1)

	actual2, err2 := parsers.ParseDriveScanMessage(input2)

	assert.NotEqual(t, nil, err2)
	assert.Equal(t, nil, actual2)

	actual3, err3 := parsers.ParseDriveScanMessage(input3)

	assert.NotEqual(t, nil, err3)
	assert.Equal(t, nil, actual3)
}
