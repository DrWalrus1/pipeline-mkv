package config

type Config struct {
	Debug     bool `json:"debug"`
	DirectIO  bool `json:"direct_io"`
	RobotMode bool `json:"robot_mode"`
}
