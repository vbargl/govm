package versionutil

import (
	"go/build"
	"strings"

	"barglvojtech.net/govm/pkg/embi/types"
	"barglvojtech.net/x/pkg/stringutil"
)

func Qualify(version types.QualifiedVersion) types.QualifiedVersion {
	return types.QualifiedVersion{
		Version: version.Version,
		Goos:    strings.ToLower(stringutil.FirstNonEmpty(version.Goos, build.Default.GOOS)),
		Goarch:  strings.ToLower(stringutil.FirstNonEmpty(version.Goos, build.Default.GOARCH)),
	}
}
