package config

import (
	"fmt"
	"slices"
)

func ValidateLLpkgConfig(config LLpkgConfig) error {
	return validateUpstreamConfig(config.UpstreamConfig)
}

func validateUpstreamConfig(config UpstreamConfig) error {
	// 1. check if upstream installer is valid
	if config.InstallerConfig.Name == "" {
		return fmt.Errorf("invalid configuration: upstream.installer.name is required")
	}
	if !slices.Contains(ValidInstallers, config.InstallerConfig.Name) {
		return fmt.Errorf("invalid configuration: upstream.installer.name %s is not supported", config.InstallerConfig.Name)
	}

	// 2. check if package is valid
	if config.PackageConfig.Name == "" {
		return fmt.Errorf("invalid configuration: upstream.package.name is required")
	}
	if config.PackageConfig.Version == "" {
		return fmt.Errorf("invalid configuration: upstream.package.version is required")
	}

	return nil
}
