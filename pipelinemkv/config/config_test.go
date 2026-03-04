package config_test

import (
	"pipelinemkv/config"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCanParseConfig(t *testing.T) {
	config, err := config.Load("./test_config.json")
	assert.Equal(t, nil, err)

	assert.Equal(t, config.Arguments.DirectIO, true)
	assert.Equal(t, config.Arguments.Cache, 1024)
	assert.Equal(t, config.Arguments.TitleMinLength, 10)
	assert.Equal(t, config.DiscReadLogLevel, "error")
	assert.Equal(t, config.LogLevel, "info")
	argsAsString := strings.Join(config.Arguments.ConvertArgumentsToArgs(), " ")

	assert.Equal(t, argsAsString, "--directio=true --minlength=10 --cache=1024")
}

func TestHasExecutablePath(t *testing.T) {
	config := config.Config{
		ExecutablePath: "../makemkvcon",
	}

	assert.Equal(t, true, config.HasAlternateExecutablePath())

	config.ExecutablePath = ""

	assert.Equal(t, false, config.HasAlternateExecutablePath())
}
