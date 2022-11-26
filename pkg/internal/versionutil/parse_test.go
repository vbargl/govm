package versionutil

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	type param struct {
		given    string
		expected VersionParts
	}

	params := []param{
		// go prefixed
		{given: "go1.20", expected: VersionParts{major: "1", minor: "20"}},
		{given: "go1.20.2", expected: VersionParts{major: "1", minor: "20", build: "2"}},
		{given: "go1.20rc3", expected: VersionParts{major: "1", minor: "20", mark: "rc3"}},
		{given: "go1.20beta4", expected: VersionParts{major: "1", minor: "20", mark: "beta4"}},
		{given: "GO1.20RC5", expected: VersionParts{major: "1", minor: "20", mark: "rc5"}},
		{given: "go1.20.6beta7", expected: VersionParts{major: "1", minor: "20", build: "6", mark: "beta7"}},
		{given: "go1.20.8rc1", expected: VersionParts{major: "1", minor: "20", build: "8", mark: "rc1"}},
		// v prefixed
		{given: "v1.20", expected: VersionParts{major: "1", minor: "20"}},
		{given: "v1.20.2", expected: VersionParts{major: "1", minor: "20", build: "2"}},
		{given: "v1.20rc3", expected: VersionParts{major: "1", minor: "20", mark: "rc3"}},
		{given: "v1.20beta4", expected: VersionParts{major: "1", minor: "20", mark: "beta4"}},
		{given: "v1.20RC5", expected: VersionParts{major: "1", minor: "20", mark: "rc5"}},
		{given: "v1.20.6beta7", expected: VersionParts{major: "1", minor: "20", build: "6", mark: "beta7"}},
		{given: "v1.20.8rc1", expected: VersionParts{major: "1", minor: "20", build: "8", mark: "rc1"}},
		// without prefix
		{given: "1.20", expected: VersionParts{major: "1", minor: "20"}},
		{given: "1.20.2", expected: VersionParts{major: "1", minor: "20", build: "2"}},
		{given: "1.20rc3", expected: VersionParts{major: "1", minor: "20", mark: "rc3"}},
		{given: "1.20beta4", expected: VersionParts{major: "1", minor: "20", mark: "beta4"}},
		{given: "1.20RC5", expected: VersionParts{major: "1", minor: "20", mark: "rc5"}},
		{given: "1.20.6beta7", expected: VersionParts{major: "1", minor: "20", build: "6", mark: "beta7"}},
		{given: "1.20.8rc1", expected: VersionParts{major: "1", minor: "20", build: "8", mark: "rc1"}},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("version %v", p.given), func(t *testing.T) {
			got, err := Parse(p.given)

			if err != nil {
				t.Fatalf("unexpected %v", err)
			}
			if got != p.expected {
				t.Errorf("expected %v, got %v", p.expected, got)
			}
		})
	}
}
