package versionutil

import (
	"fmt"

	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/govm/pkg/internal/archiveutil"

	"barglvojtech.net/x/pkg/stringutil"
)

const (
	// FilenameFormat is used for naming.
	FilenameFormat = "${version}.${goos}-${goarch}.${ext}"
	// DefaultSdkUrl is url where all versions of go are available
	UrlFormat = "https://go.dev/dl/${version}.${goos}-${goarch}.${ext}"
)

// Format will format QualifiedVersion using format variable.
// This function will interpolate variables (version, goos, goarch, ext) using [barglvojtech.net/x/pkg/stringutil.Interpolate].
func Format(format string, version types.QualifiedVersion) (string, error) {
	if format == "" {
		return "", fmt.Errorf("versionutil.Format: format is required")
	}
	if version.Version == "" {
		return "", fmt.Errorf("versionutil.Format: version.Version is required")
	}

	version = QualifyDefaults(version)
	format = stringutil.FirstNonEmpty(format, FilenameFormat)

	return stringutil.Interpolate(format, map[string]string{
		"version": version.Version.String(),
		"goos":    version.Goos,
		"goarch":  version.Goarch,
		"ext":     archiveutil.GetExtension(version.Goos),
	}), nil
}
