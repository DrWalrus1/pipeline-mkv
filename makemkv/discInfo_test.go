package makemkv_test

import (
	"servermakemkv/makemkv"
	"servermakemkv/makemkv/commands/outputs"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUpdateDiscInfo(t *testing.T) {
	name := "Demon Slayer"
	discType := "Blu-ray disc"
	language := "English"
	expected := makemkv.DiscInfo{
		Name:     name,
		Type:     discType,
		Language: language,
	}

	var actual makemkv.DiscInfo

	typeDiscInfo := outputs.DiscInformation{
		ID:    outputs.Type,
		Value: discType,
	}
	nameDiscInfo := outputs.DiscInformation{
		ID:    outputs.Name,
		Value: name,
	}
	languageDiscInfo := outputs.DiscInformation{
		ID:    outputs.MetadataLanguageName,
		Value: language,
	}

	actual.UpdateDiscInfo(typeDiscInfo)
	actual.UpdateDiscInfo(nameDiscInfo)
	actual.UpdateDiscInfo(languageDiscInfo)

	assert.Equal(t, expected, actual)
}

func TestUpdateTitle(t *testing.T) {
	// name := "Demon Slayer"
	// size := "29.5GB"
	// sizeInBytes := "295123098"
	// duration := "5 hours"
	// language := "English"
	// chapters := "24"
	// outputfilename := "test.txt"
	//
	// expectedTitleInfo := makemkv.Title{}

}
