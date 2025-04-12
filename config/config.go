package config

type Config struct {
	ExecutablePath string `json:"executable_path"`
	Arguments      struct {
		Debug     bool `json:"debug"`
		DirectIO  bool `json:"direct_io"`
		RobotMode bool `json:"robot_mode"`
	}
}

func (c *Config) IntoCLIArguments() string {
	return ""
}
