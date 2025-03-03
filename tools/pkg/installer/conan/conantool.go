package conan

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/goplus/llpkg/tools/internal/cmdbuilder"
	"github.com/goplus/llpkg/tools/pkg/config"
	"github.com/goplus/llpkg/tools/pkg/installer"
)

const (
	ConanfileTemplate = `[requires]
	%s/%s

	[generators]
	PkgConfigDeps

	[options]
	*:shared=True
	%s`
)

type conanInstaller struct {
	pkg    config.Package
	config map[string]string
}

func NewConanInstaller(pkg config.Package, config map[string]string) installer.Installer {
	return &conanInstaller{
		pkg:    pkg,
		config: config,
	}
}

func (c *conanInstaller) Name() string {
	return "conan"
}

func (c *conanInstaller) Config() map[string]string {
	return c.config
}

func (c *conanInstaller) Install(dir string) error {
	return c.install(dir)
}

// String prints the specified conanfile
func (c *conanInstaller) String() string {
	return fmt.Sprintf(ConanfileTemplate, c.pkg.Name, c.pkg.Version, c.pkg.Name)
}

func (c *conanInstaller) options() string {
	return strings.Join([]string{"\\*:shared=True", c.config["options"]}, " ")
}

func (c *conanInstaller) install(dir string) error {
	// Build the following command
	// conan install --requires %s -g PkgConfigDeps --options \\*:shared=True --build=missing --output-folder=%s\
	builder := cmdbuilder.NewCmdBuilder(cmdbuilder.WithConanSerilazier())
	builder.Set("requires", c.pkg.Name+"/"+c.pkg.Version)
	builder.Set("generator", "PkgConfigDeps")
	builder.Set("options", c.options())
	builder.Set("build", "missing")
	builder.Set("output-folder", dir)

	args := append([]string{"install"}, builder.Args()...)

	buildCmd := exec.Command("conan", args...)

	fmt.Println(buildCmd)

	out, err := buildCmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}
