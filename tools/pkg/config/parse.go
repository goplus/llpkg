package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseLLpkgConfig(configPath string) (LLpkgConfig, error) {
	var config LLpkgConfig
	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, fmt.Errorf("failed to decode config file: %w", err)
	}

	// set default values
	config = fillDefaults(config)

	return config, nil
}

func fillDefaults(config LLpkgConfig) LLpkgConfig {
	if config.UpstreamConfig.InstallerConfig.Name == "" {
		config.UpstreamConfig.InstallerConfig.Name = "conan"
	}
	return config
}
