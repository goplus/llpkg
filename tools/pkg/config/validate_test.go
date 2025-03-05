package config

import "testing"

func TestValidateLLpkgConfig(t *testing.T) {
	config, err := ParseLLpkgConfig("../../_demo/llpkg.cfg")
	if err != nil {
		t.Errorf("Error parsing config file: %v", err)
	}
	err = ValidateLLpkgConfig(config)
	if err != nil {
		t.Errorf("Error validating config: %v", err)
	}
}
