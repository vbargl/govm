package versionutil

import (
	"strings"

	"barglvojtech.net/govm/pkg/embi/types"
)

func IsBeta(version types.Version) bool {
	return strings.Contains(version.String(), "beta")
}

func IsReleaseCandidate(version types.Version) bool {
	return strings.Contains(version.String(), "rc")
}

func isStable(version types.Version) bool {
	return !IsBeta(version) && !IsReleaseCandidate(version)
}
