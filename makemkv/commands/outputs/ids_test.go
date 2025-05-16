package outputs

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConvertItemAttributeIntoToString(t *testing.T) {
	tests := map[string]int{
		`Unknown`:                            0,
		"Type":                               1,
		"Name":                               2,
		"LanguageCode":                       3,
		"LanguageName":                       4,
		"CodecID":                            5,
		"ShortCodecName":                     6,
		"LongCodecName":                      7,
		"NumberOfChapters":                   8,
		"Duration":                           9,
		"DiskSize":                           10,
		"DiskSizeInBytes":                    11,
		"StreamTypeExtension":                12,
		"Bitrate":                            13,
		"NumberOfAudioChannels":              14,
		"AngleInformation":                   15,
		"Comment":                            49,
		"OffsetSequenceID":                   50,
		"MaxValue":                           51,
		"Unknown Application Item Attribute": 60,
	}
	for expected, tt := range tests {
		t.Run(expected, func(t *testing.T) {
			id, err := GetItemAttributeDescription(tt)
			if err != nil {
				assert.Equal(t, expected, err.Error())
			} else {
				assert.Equal(t, expected, id)
			}
		})

	}
}
