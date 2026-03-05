package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	// Executable path for MakeMkv
	ExecutablePath  string `json:"executable_path"`
	RegistrationKey string `json:"registration_key"`
	// Specifies how much read logs from MakeMKV get sent back to the client.
	DiscReadLogLevel string `json:"disc_read_log_level"`
	// Specifies what level of logs from the server get logged
	LogLevel             string `json:"log_level"`
	Arguments            Arguments
	MetadataServiceToken string `json:"metadata_service_token"`
	Port                 string
}

func (c *Config) HasAlternateExecutablePath() bool {
	return c.ExecutablePath != ""
}

func (c *Config) ConvertConfigToArgs() []string {
	args := []string{}

	if c.ExecutablePath != "" {
		args = append(args, c.ExecutablePath)
	}

	var requiredMakeMkvOptions = []string{"-r", "--progress=-stdout"}
	// We always enable robot mode
	args = append(args, requiredMakeMkvOptions...)

	args = append(args, c.Arguments.ConvertArgumentsToArgs()...)

	return args
}

func Load(flagPath string) (*Config, error) {
	path, err := resolvePath(flagPath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config at %s: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config at %s: %w", path, err)
	}

	cfg.Port = getPort()
	//TODO: Add error handling

	return &cfg, nil
}

func resolvePath(flagPath string) (string, error) {
	candidates := buildConfigCandidates(flagPath)

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("no config file found; tried: %v", candidates)
}

func buildConfigCandidates(flagPath string) []string {
	// var configPath string
	// flag.StringVar(&configPath, "config", "", "filepath for config.json file")
	if flagPath != "" {
		return []string{flagPath}
	}

	var candidates []string

	// 2. Environment variable
	if env := os.Getenv("PIPELINEMKV_CONFIG"); env != "" {
		candidates = append(candidates, env)
		return candidates // same — explicit, don't fall through
	}

	// 3. User-local (XDG)
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		if home, err := os.UserHomeDir(); err == nil {
			configDir = filepath.Join(home, ".config")
		}
	}
	if configDir != "" {
		candidates = append(candidates, filepath.Join(configDir, "pipelinemkv", "config.json"))
	}

	// 4. System-wide fallback
	candidates = append(candidates, "/etc/pipelinemkv/config.json")

	return candidates
}

func getPort() string {
	var port string
	flag.StringVar(&port, "port", "", "Port to host the server on")
	flag.Parse()

	if port != "" {
		return fmt.Sprintf(":%s", port)
	}

	// 2. Environment variable
	if env := os.Getenv("PIPELINEMKV_CONFIG"); env != "" {
		return fmt.Sprintf(":%s", env)
	}

	return ":9090"
}
