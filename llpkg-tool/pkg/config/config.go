package config

type LLpkgConfig struct {
	Package   Package   `json:"package"`
	Upstream  Upstream  `json:"upstream"`
	ToolChain ToolChain `json:"toolchain"`
}

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Upstream struct {
	Name   string            `json:"name"`
	Config map[string]string `json:"config"`
}

type ToolChain struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Config  map[string]string `json:"config"`
}
