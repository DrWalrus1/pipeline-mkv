package config

import (
	"fmt"
)

type Arguments struct {
	// Enables or disables direct disc access
	DirectIO bool `json:"direct_io"`
	// Default is 120 seconds
	TitleMinLength int `json:"title_min_length"`
	// Specifies size of read cache in megabytes used by MakeMKV. By default program uses huge amount of memory. About 128 MB is recommended for streaming and backup, 512MB for DVD conversion and 1024MB for Blu-ray conversion.
	Cache int `json:"cache"`
}

func (a *Arguments) ConvertArgumentsToArgs() []string {
	args := []string{}

	if a.DirectIO {
		args = append(args, fmt.Sprintf("--directio=true"))
	}
	if a.TitleMinLength > -1 {
		args = append(args, fmt.Sprintf("--minlength=%d", a.TitleMinLength))
	}
	if a.Cache > 0 {
		args = append(args, fmt.Sprintf("--cache=%d", a.Cache))
	}
	return args
}

type Config struct {
	ExecutablePath  string `json:"executable_path"`
	RegistrationKey string `json:"registration_key"`
	Arguments       Arguments
}

func (c *Config) ConvertConfigToArgs() []string {
	args := []string{}

	if c.ExecutablePath != "" {
		args = append(args, c.ExecutablePath)
	}

	// We always enable robot mode
	args = append(args, "-r")

	args = append(args, c.Arguments.ConvertArgumentsToArgs()...)

	return args
}
