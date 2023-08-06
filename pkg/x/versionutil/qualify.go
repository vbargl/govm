package versionutil

import (
	"go/build"
	"strings"

	"barglvojtech.net/govm/pkg/embi/types"

	"barglvojtech.net/x/pkg/stringutil"
)

func Qualify(version types.Version) types.QualifiedVersion {
	return types.QualifiedVersion{
		Version: version,
		Goos:    strings.ToLower(build.Default.GOOS),
		Goarch:  strings.ToLower(build.Default.GOARCH),
	}
}

func QualifyDefaults(version types.QualifiedVersion) types.QualifiedVersion {
	return types.QualifiedVersion{
		Version: version.Version,
		Goos:    strings.ToLower(stringutil.FirstNonEmpty(version.Goos, build.Default.GOOS)),
		Goarch:  strings.ToLower(stringutil.FirstNonEmpty(version.Goarch, build.Default.GOARCH)),
	}
}
