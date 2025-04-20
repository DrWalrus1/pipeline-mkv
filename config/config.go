package config

import (
	"fmt"
)

type Arguments struct {
	DirectIO       bool `json:"direct_io"`
	TitleMinLength int  `json:"title_min_length"`
	Cache          int  `json:"cache"` // TODO: Should provide the suggestions from the usage file so end users to set this incorrectly
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
