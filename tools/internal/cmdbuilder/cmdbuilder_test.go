package cmdbuilder

import (
	"slices"
	"testing"
)

func TestCmdBuilder(t *testing.T) {
	conanBuilder := NewCmdBuilder(WithConanSerilazier())
	conanBuilder.Set("options", "\\*:shared=True")
	conanBuilder.Set("build", "missing")
	if !slices.Equal(conanBuilder.Args(), []string{"--options=\\*:shared=True", "--build=missing"}) {
		t.Errorf("Unexpected args: %v", conanBuilder.Args())
	}
	if conanBuilder.String() != "--options=\\*:shared=True --build=missing" {
		t.Errorf("Unexpected string: %s", conanBuilder.String())
	}
}
