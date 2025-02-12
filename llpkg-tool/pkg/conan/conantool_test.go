package conan

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/goplus/llpkg/llpkg-tool/pkg/config"
)

func TestConanTool(t *testing.T) {
	conanTool := &ConanTool{
		Package: config.Package{
			Name:    "cjson",
			Version: "1.7.18",
		},
	}

	// 1. simple run
	err := conanTool.refreshWorkingDir()
	if err != nil {
		t.Errorf("Error refreshing working dir: %v", err)
	}

	err = conanTool.Run()
	if err != nil {
		t.Errorf("Error running conan tool: %v", err)
	}

	err = verify(conanTool)
	if err != nil {
		t.Errorf("Error verifying conan tool: %v", err)
	}

	// 2. generate conanfile.txt
	err = conanTool.refreshWorkingDir()
	if err != nil {
		t.Errorf("Error refreshing working dir: %v", err)
	}

	err = conanTool.genConanFile()
	if err != nil {
		t.Errorf("Error generating conanfile.txt: %v", err)
	}

	err = verifyConanFile(conanTool)
	if err != nil {
		t.Errorf("Error verifying conanfile.txt: %v", err)
	}

	// 3. install

	err = conanTool.install()
	if err != nil {
		t.Errorf("Error installing: %v", err)
	}

	err = verifyConanInstall(conanTool)
	if err != nil {
		t.Errorf("Error verifying installation: %v", err)
	}
}

func mkdirTemp() (string, error) {
	dir, err := os.MkdirTemp("", "llpkg-tool")
	if err != nil {
		return "", err
	}
	return dir, nil
}

func (c *ConanTool) refreshWorkingDir() (err error) {
	c.WorkingDir, err = mkdirTemp()
	fmt.Println("Working dir: " + c.WorkingDir)
	return err
}

func verify(conanTool *ConanTool) error {
	err := verifyConanFile(conanTool)
	if err != nil {
		return err
	}

	err = verifyConanInstall(conanTool)
	if err != nil {
		return err
	}

	return nil
}

func verifyConanFile(conanTool *ConanTool) error {
	conanFile, err := os.Open(conanTool.WorkingDir + "/conanfile.txt")
	if err != nil {
		return err
	}

	fileContent := make([]byte, 1024)
	len, err := conanFile.Read([]byte(fileContent))
	if err != nil {
		return err
	}

	if string(fileContent)[0:len] != conanTool.toString() {
		return errors.New("File content is not as expected, expected:\n --- \n " + conanTool.toString() + "\n --- \nactual: \n --- \n" + string(fileContent) + "\n --- \n")
	}

	return nil
}

func verifyConanInstall(conanTool *ConanTool) error {
	// 1. ensure .pc file exists
	_, err := os.Stat(conanTool.WorkingDir + "/" + conanTool.Package.Name + ".pc")
	if err != nil {
		return errors.New(".pc file does not exist: " + err.Error())
	}

	// 2. ensure pkg-config can find .pc file
	os.Setenv("PKG_CONFIG_PATH", conanTool.WorkingDir)
	defer os.Unsetenv("PKG_CONFIG_PATH")

	buildCmd := exec.Command("pkg-config", "--cflags", conanTool.Package.Name)
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		return errors.New("pkg-config failed: " + err.Error() + " with output: " + string(out))
	}

	return nil
}
