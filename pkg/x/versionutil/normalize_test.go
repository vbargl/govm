package versionutil

import (
	"fmt"
	"testing"

	"barglvojtech.net/govm/pkg/embi/types"
)

func TestNormalizeCorrect(t *testing.T) {
	type param struct {
		given    string
		expected types.Version
	}

	params := []param{
		// go prefixed
		{given: "go1.20", expected: "go1.20"},
		{given: "go1.20.2", expected: "go1.20.2"},
		{given: "go1.20rc3", expected: "go1.20rc3"},
		{given: "go1.20beta4", expected: "go1.20beta4"},
		{given: "GO1.20RC5", expected: "go1.20rc5"},
		// v prefixed
		{given: "v1.20", expected: "go1.20"},
		{given: "v1.20.2", expected: "go1.20.2"},
		{given: "v1.20rc3", expected: "go1.20rc3"},
		{given: "v1.20beta4", expected: "go1.20beta4"},
		{given: "V1.20RC5", expected: "go1.20rc5"},
		// without prefix
		{given: "1.20", expected: "go1.20"},
		{given: "1.20.2", expected: "go1.20.2"},
		{given: "1.20rc3", expected: "go1.20rc3"},
		{given: "1.20beta4", expected: "go1.20beta4"},
		{given: "1.20RC5", expected: "go1.20rc5"},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("version %v", p.given), func(t *testing.T) {
			got, err := Normalize(p.given)

			if err != nil {
				t.Fatalf("unexpected %v", err)
			}
			if got != p.expected {
				t.Errorf("expected %v, got %v", p.expected, got)
			}
		})
	}
}

func TestNormalizeMalformated(t *testing.T) {

	params := []string{
		"",
		"something else",
		"1.20rc",   // rc without number
		"1.30beta", // beta without number,
		"1.40something",
	}

	for _, s := range params {
		t.Run(fmt.Sprintf("version %v", s), func(t *testing.T) {
			if ver, err := Normalize(s); err == nil {
				t.Errorf("expected error for %s returned %s, %v", s, ver, err)
			}
		})
	}
}
