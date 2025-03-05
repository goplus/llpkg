package config

var ValidInstallers = []string{"conan"}

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
