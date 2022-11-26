package opsutil

import (
	"fmt"
	"go/build"
	"strings"

	"barglvojtech.net/govm/pkg/embi/types"
)

const (
	fileExtensionWindows = "zip"
	fileExtensionUnix    = "tar.gz"
)

func VersionFile(sdkUrl string, version types.Version) types.VersionFile {
	versionFile := types.VersionFile{
		Version:  version,
		Url:      "",
		Goarch:   build.Default.GOARCH,
		Goos:     build.Default.GOOS,
		Checksum: "",
	}

	versionFile.Url = interpolate(sdkUrl, map[string]string{
		"version": versionFile.Version.String(),
		"goos":    versionFile.Goos,
		"goarch":  versionFile.Goarch,
		"ext":     getExt(),
	})

	return versionFile
}

func getExt() string {
	switch build.Default.GOOS {
	case "windows":
		return fileExtensionWindows
	default:
		return fileExtensionUnix
	}
}

func interpolate(s string, variables map[string]string) string {
	for variable, value := range variables {
		s = strings.ReplaceAll(s, fmt.Sprintf("${%s}", variable), value)
	}
	return s
}
