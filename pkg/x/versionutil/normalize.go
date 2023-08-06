package versionutil

import (
	"barglvojtech.net/govm/pkg/embi/types"
)

// Normalize will parse multiple formats into official go version.
// The output format will be go1.2.3 (can contain beta and rc).
// Returns version or error if could not be parsed.
//
// See test to know which cases are handled and which are not.
func Normalize(version string) (types.Version, error) {
	parts, err := Parse(version)
	if err != nil {
		return "", err
	}
	return parts.Version(), err
}
