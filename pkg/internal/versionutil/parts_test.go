package versionutil

import (
	"testing"

	"barglvojtech.net/govm/pkg/embi/types"
)

func TestVersionPartsToVersion(t *testing.T) {
	type param struct {
		given    VersionParts
		expected types.Version
	}

	params := []param{
		{given: VersionParts{major: "1", minor: "20", build: "", mark: ""}, expected: "go1.20"},
		{given: VersionParts{major: "1", minor: "20", build: "", mark: "rc1"}, expected: "go1.20rc1"},
		{given: VersionParts{major: "1", minor: "20", build: "3", mark: ""}, expected: "go1.20.3"},
		{given: VersionParts{major: "1", minor: "20", build: "3", mark: "beta1"}, expected: "go1.20.3beta1"},
	}

	for _, p := range params {
		if ver := p.given.Version(); ver != p.expected {
			t.Errorf("expected %s, got %s", p.expected, ver)
		}
	}
}
