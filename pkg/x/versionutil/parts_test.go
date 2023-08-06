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
		{given: VersionParts{Major: "1", Minor: "20", Patch: "", PreRelease: ""}, expected: "go1.20"},
		{given: VersionParts{Major: "1", Minor: "20", Patch: "", PreRelease: "rc1"}, expected: "go1.20rc1"},
		{given: VersionParts{Major: "1", Minor: "20", Patch: "3", PreRelease: ""}, expected: "go1.20.3"},
		{given: VersionParts{Major: "1", Minor: "20", Patch: "3", PreRelease: "beta1"}, expected: "go1.20.3beta1"},
	}

	for _, p := range params {
		if ver := p.given.Version(); ver != p.expected {
			t.Errorf("expected %s, got %s", p.expected, ver)
		}
	}
}

func TestIsBeta(t *testing.T) {
	type param struct {
		given  types.Version
		isBeta bool
	}

	params := []param{
		{given: "1.20beta1", isBeta: true},
		{given: "1.20.2beta1", isBeta: true},
		{given: "1.20rc1", isBeta: false},
		{given: "1.20.1", isBeta: false},
		{given: "1.20", isBeta: false},
	}

	for _, p := range params {
		isBeta := IsBeta(p.given)
		if isBeta != p.isBeta {
			t.Errorf("expected isBeta %v = %v, got %v", p.given, p.isBeta, isBeta)
		}
	}
}

func TestIsReleaseCandidate(t *testing.T) {
	type param struct {
		given              types.Version
		isReleaseCandidate bool
	}

	params := []param{
		{given: "1.20rc1", isReleaseCandidate: true},
		{given: "1.20.2rc1", isReleaseCandidate: true},
		{given: "1.20beta1", isReleaseCandidate: false},
		{given: "1.20.1", isReleaseCandidate: false},
		{given: "1.20", isReleaseCandidate: false},
	}

	for _, p := range params {
		isReleaseCandidate := IsReleaseCandidate(p.given)
		if isReleaseCandidate != p.isReleaseCandidate {
			t.Errorf("expected isReleaseCandidate %v = %v, got %v", p.given, p.isReleaseCandidate, isReleaseCandidate)
		}
	}
}

func TestIsStable(t *testing.T) {
	type param struct {
		given    types.Version
		isStable bool
	}

	params := []param{
		{given: "1.20rc1", isStable: false},
		{given: "1.20.2rc1", isStable: false},
		{given: "1.20beta1", isStable: false},
		{given: "1.20.2beta1", isStable: false},
		{given: "1.20.1", isStable: true},
		{given: "1.20", isStable: true},
	}

	for _, p := range params {
		isStable := IsStable(p.given)
		if isStable != p.isStable {
			t.Errorf("expected isStable %v = %v, got %v", p.given, p.isStable, isStable)
		}
	}
}
