package config_test

import (
	"encoding/json"
	"servermakemkv/config"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCanParseConfig(t *testing.T) {
	input :=
		`
	{
		"arguments": {
			"debug": true,
			"direct_io": true,
			"robot_mode": true
		}
	}
	`
	var config config.Config

	json.Unmarshal([]byte(input), &config)

	assert.Equal(t, config.Arguments.Debug, true)
	assert.Equal(t, config.Arguments.DirectIO, true)
	assert.Equal(t, config.Arguments.RobotMode, true)
}
