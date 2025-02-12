package version

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
)

// Convert any C version into SemVer
func ToSemVer(ver string) (*semver.Version, error) {
	if strings.Trim(ver, " ") == "" {
		return nil, errors.New("empty version")
	}
	//At least two parts, "2","v1","20230607" would not be considered as SemVer
	twoPartVer := regexp.MustCompile(`^v?(0|[1-9]\d*)(?:\.(0|[1-9]\d*))(?:\.(0|[1-9]\d*))?(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if !twoPartVer.MatchString(ver) {
		return insertVer(ver)
	}
	//Convert some “looks like SemVer” version numbers (v1.2.0, 1.3, etc) into SemVer
	v, err := semver.NewVersion(ver)
	if err != nil {
		return insertVer(ver)
	}
	return v, nil
}

// Add orginal version that is not SemVer to the pre-release part of "1.0.0"
func insertVer(ver string) (*semver.Version, error) {
	//Before add to the pre-release part, replace "." with "-"
	newVer := fmt.Sprintf("0.0.0-0-%s", strings.ReplaceAll(ver, ".", "-"))
	version, err := semver.StrictNewVersion(newVer)
	if err != nil {
		return nil, errors.New("fail to convert " + ver)
	} else {
		return version, nil
	}
}
