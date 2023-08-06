package versionutil

import (
	"fmt"
	"strings"

	"barglvojtech.net/govm/pkg/internal/regexputil"
)

func Parse(version string) (p VersionParts, _ error) {
	version = strings.ToLower(version)
	version = strings.TrimPrefix(version, "go")
	version = strings.TrimPrefix(version, "v")

	var parsed int

	switch f, _ := regexputil.FindString(`\d+`, version[parsed:]); {
	case f != "":
		p.Major = f
		parsed += len(f)
	default:
		return VersionParts{}, fmt.Errorf("major version is required")
	}

	switch f, _ := regexputil.FindString(`\.\d+`, version[parsed:]); {
	case f != "":
		p.Minor = f[1:]
		parsed += len(f)
	default:
		return VersionParts{}, fmt.Errorf("minor version is required")
	}

	if f, _ := regexputil.FindString(`\.\d+`, version[parsed:]); f != "" {
		p.Patch = f[1:]
		parsed += len(f)
	}

	if f, _ := regexputil.FindString(`rc\d+`, version[parsed:]); p.PreRelease == "" && f != "" {
		p.PreRelease = f
		parsed += len(f)
	}

	if f, _ := regexputil.FindString(`beta\d+`, version[parsed:]); p.PreRelease == "" && f != "" {
		p.PreRelease = f
		parsed += len(f)
	}

	if len(version) != parsed {
		return VersionParts{}, fmt.Errorf("malformed version")
	}

	return p, nil
}
