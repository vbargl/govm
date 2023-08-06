package versionutil

import (
	"strings"

	"barglvojtech.net/govm/pkg/embi/types"
)

type VersionParts struct {
	Major      string
	Minor      string
	Patch      string
	PreRelease string
}

func (p *VersionParts) Version() types.Version {
	var b strings.Builder
	b.WriteString("go")
	b.WriteString(p.Major)
	b.WriteRune('.')
	b.WriteString(p.Minor)
	if p.Patch != "" {
		b.WriteRune('.')
		b.WriteString(p.Patch)
	}
	if p.PreRelease != "" {
		b.WriteString(p.PreRelease)
	}
	return types.Version(b.String())
}

// IsBeta checks if version contains 'beta'.
func IsBeta(version types.Version) bool {
	parts, err := Parse(version.String())
	if err != nil {
		return false
	}

	return isBeta(parts)
}

// IsReleaseCandidate checks if version contains 'rc'.
func IsReleaseCandidate(version types.Version) bool {
	parts, err := Parse(version.String())
	if err != nil {
		return false
	}

	return isReleaseCandidate(parts)
}

// IsStable checks if version does not contain 'beta' neither 'rc'.
func IsStable(version types.Version) bool {
	parts, err := Parse(version.String())
	if err != nil {
		return false
	}

	return !isBeta(parts) && !isReleaseCandidate(parts)
}

func isReleaseCandidate(vp VersionParts) bool {
	return strings.HasPrefix(vp.PreRelease, "rc")
}

func isBeta(vp VersionParts) bool {
	return strings.HasPrefix(vp.PreRelease, "beta")
}
