package tools

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	sem "github.com/Masterminds/semver/v3"
)

// Convert any C version into SemVer
func ToSemVer(ver string) (*sem.Version, error) {
	//At least two parts, "2","v1","20230607" are not allowed
	twoPartVer := regexp.MustCompile(`^v?(0|[1-9]\d*)(?:\.(0|[1-9]\d*))(?:\.(0|[1-9]\d*))?(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if !twoPartVer.MatchString(ver) {
		return insertVer(ver)
	}
	//Convert some “looks like SemVer” version numbers (v1.2.0, 1.3, etc) into SemVer
	v, err := sem.NewVersion(ver)
	if err != nil {
		return insertVer(ver)
	}
	return v, nil
}

// Add orginal version that is not SemVer to the pre-release part of "1.0.0"
func insertVer(ver string) (*sem.Version, error) {
	metaVersion := regexp.MustCompile(`^.+(\+).+$`)
	var newVer string
	//Before add to the pre-release part, add "-llgo" suffix before "+" (if exists), and replace "." with "-"
	if metaVersion.MatchString(ver) {
		newVer = strings.Replace(ver, "+", "-llgo+", 1)
		newVer = fmt.Sprintf("1.0.0-%s", strings.ReplaceAll(newVer, ".", "-"))
	} else {
		newVer = fmt.Sprintf("1.0.0-%s-llgo", strings.ReplaceAll(ver, ".", "-"))
	}
	version, err := sem.StrictNewVersion(newVer)
	if err != nil {
		return nil, errors.New("Fail to convert " + ver)
	} else {
		return version, nil
	}
}
