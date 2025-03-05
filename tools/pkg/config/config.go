package config

import (
	"fmt"

	"github.com/goplus/llpkg/tools/pkg/upstream"
	"github.com/goplus/llpkg/tools/pkg/upstream/installer/conan"
)

type LLpkgConfig struct {
	UpstreamConfig UpstreamConfig `json:"upstream"`
}

type UpstreamConfig struct {
	InstallerConfig InstallerConfig `json:"installer"`
	PackageConfig   PackageConfig   `json:"package"`
}

type InstallerConfig struct {
	Name   string            `json:"name"`
	Config map[string]string `json:"config,omitempty"`
}

type PackageConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewUpstreamFromConfig(config UpstreamConfig) (upstream.Upstream, error) {
	var installer upstream.Installer
	switch config.InstallerConfig.Name {
	case "conan":
		installer = conan.NewConanInstaller(config.InstallerConfig.Config)
	default:
		return nil, fmt.Errorf("invalid configuration: upstream.installer.name %s is not supported", config.InstallerConfig.Name)
	}

	var up upstream.Upstream = upstream.NewDefaultUpstream(installer, upstream.Package{
		Name:    config.PackageConfig.Name,
		Version: config.PackageConfig.Version,
	})

	return up, nil
}
