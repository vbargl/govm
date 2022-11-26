package versionutil

import (
	"fmt"

	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/govm/pkg/internal/archiveutil"
	"barglvojtech.net/x/pkg/stringutil"
)

const (
	// DefaultSdkUrl is url where all versions of go are available
	DefaultVersionUrl = "https://go.dev/dl/${version}.${goos}-${goarch}.${ext}"
)

func Url(format string, version types.QualifiedVersion) (string, error) {
	if version.Version == "" {
		return "", fmt.Errorf("version is required param")
	}

	version = Qualify(version)
	format = stringutil.FirstNonEmpty(format, DefaultVersionUrl)

	return interpolate(format, map[string]string{
		"version": version.Version.String(),
		"goos":    version.Goos,
		"goarch":  version.Goarch,
		"ext":     archiveutil.GetExtension(version.Goos),
	}), nil
}
