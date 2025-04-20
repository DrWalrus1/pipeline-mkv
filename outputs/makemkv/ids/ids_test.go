package ids_test

import (
	"servermakemkv/outputs/makemkv/ids"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConvertItemAttributeIntoToString(t *testing.T) {
	tests := map[string]uint{
		`Unknown`:                            0,
		"Type":                               1,
		"Name":                               2,
		"Language Code":                      3,
		"Language Name":                      4,
		"Codec ID":                           5,
		"Short Codec Name":                   6,
		"Long Codec Name":                    7,
		"Number of Chapters":                 8,
		"Duration":                           9,
		"Disk Size":                          10,
		"Disk Size in Bytes":                 11,
		"Stream Type Extension":              12,
		"Bitrate":                            13,
		"Number of Audio Channels":           14,
		"Angle Information":                  15,
		"Comment":                            49,
		"Offset Sequence ID":                 50,
		"Maximum Value":                      51,
		"Unknown Application Item Attribute": 60,
	}
	for expected, tt := range tests {
		t.Run(expected, func(t *testing.T) {
			assert.Equal(t, expected, ids.GetItemAttributeDescription(tt))
		})
	}
}
