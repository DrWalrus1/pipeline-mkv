package config_test

import (
	"servermakemkv/config"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCanParseConfig(t *testing.T) {
	input :=
		`
	{
		"arguments": {
			"direct_io": true,
			"cache": 1024,
			"title_min_length": 10
		}
	}
	`
	reader := strings.NewReader(input)

	config, err := config.LoadConfig(reader)
	assert.Equal(t, nil, err)

	assert.Equal(t, config.Arguments.DirectIO, true)
	assert.Equal(t, config.Arguments.Cache, 1024)
	assert.Equal(t, config.Arguments.TitleMinLength, 10)
	argsAsString := strings.Join(config.Arguments.ConvertArgumentsToArgs(), " ")

	assert.Equal(t, argsAsString, "--directio=true --minlength=10 --cache=1024")
}
