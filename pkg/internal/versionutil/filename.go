package versionutil

import (
	"fmt"

	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/govm/pkg/internal/archiveutil"
	"barglvojtech.net/x/pkg/stringutil"
)

const (
	// DefaultVersionFilename is used for naming.
	// Keep same as DefaultVersionUrl
	DefaultVersionFilename = "${version}.${goos}-${goarch}.${ext}"
)

// Filename
func Filename(format string, version types.QualifiedVersion) (string, error) {
	if version.Version == "" {
		return "", fmt.Errorf("version is required param")
	}

	version = Qualify(version)
	format = stringutil.FirstNonEmpty(format, DefaultVersionFilename)

	return stringutil.Interpolate(format, map[string]string{
		"version": version.Version.String(),
		"goos":    version.Goos,
		"goarch":  version.Goarch,
		"ext":     archiveutil.GetExtension(version.Goos),
	}), nil
}
