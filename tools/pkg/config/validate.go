package config

import (
	"fmt"
	"slices"

	"github.com/goplus/llpkg/tools/pkg/upstream"
)

func ValidateLLpkgConfig(config LLpkgConfig) error {
	return validateConfigUpstream(config)
}

func validateConfigUpstream(config LLpkgConfig) error {
	// 1. check if upstream installer is valid
	if config.UpstreamConfig.InstallerConfig.Name == "" {
		return fmt.Errorf("invalid configuration: upstream.installer.name is required")
	}
	if !slices.Contains(upstream.ValidInstallers, config.UpstreamConfig.InstallerConfig.Name) {
		return fmt.Errorf("invalid configuration: upstream.installer.name %s is not supported", config.UpstreamConfig.InstallerConfig.Name)
	}

	// 2. check if package is valid
	if config.UpstreamConfig.PackageConfig.Name == "" {
		return fmt.Errorf("invalid configuration: upstream.package.name is required")
	}
	if config.UpstreamConfig.PackageConfig.Version == "" {
		return fmt.Errorf("invalid configuration: upstream.package.version is required")
	}

	// 3. build upstream, and check if package exists
	up, err := NewUpstreamFromConfig(config.UpstreamConfig)
	if err != nil {
		return fmt.Errorf("failed to build upstream: %w", err)
	}
	_, err = up.Installer().Search(up.Package())
	if err != nil {
		return fmt.Errorf("failed to search for package %s: %w", up.Package().Name, err)
	}

	return nil
}
