package conan

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/goplus/llpkg/llpkg-tool/pkg/config"
)

type ConanTool struct {
	Package    config.Package
	Config     map[string]string
	WorkingDir string
}

func (c *ConanTool) Run() error {
	err := c.genConanFile()
	if err != nil {
		return err
	}

	err = c.install()
	if err != nil {
		return err
	}

	return nil
}

func (c *ConanTool) genConanFile() error {
	err := os.WriteFile(c.WorkingDir+"/conanfile.txt", []byte(c.toString()), 0o644)
	return err
}

func (c *ConanTool) toString() string {
	return "[requires]\n" + c.Package.Name + "/" + c.Package.Version + "\n\n" +
		"[generators]\nPkgConfigDeps\n\n" +
		"[options]\n" + c.Config["options"]
}

func (c *ConanTool) install() error {
	buildCmd := exec.Command("conan", "install", c.WorkingDir, "--build=missing", "-o \\*:shared=True", "--output-folder="+c.WorkingDir)
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}
