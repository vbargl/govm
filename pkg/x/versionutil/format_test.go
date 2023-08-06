package versionutil

import (
	"fmt"
	"testing"

	"barglvojtech.net/govm/pkg/embi/types"
)

func TestFormat(t *testing.T) {
	type param struct {
		// given
		format  string
		version types.QualifiedVersion

		// expected
		filename string
		err      error
	}

	params := []param{
		{
			format: "${version}/${goos}/${goarch}.${ext}",
			version: types.QualifiedVersion{
				Version: "go1.20",
				Goos:    "windows",
				Goarch:  "arm64",
			},

			filename: "go1.20/windows/arm64.zip",
		},
		{
			format: "${version}@${goos}-${goarch}.${ext}",
			version: types.QualifiedVersion{
				Version: "go1.20.3",
				Goos:    "linux",
				Goarch:  "arm64",
			},

			filename: "go1.20.3@linux-arm64.tar.gz",
		},
		{
			format: "${version}#${goos}|${goarch}@${ext}",
			version: types.QualifiedVersion{
				Version: "go1.20.3",
				Goos:    "macos",
				Goarch:  "arm32",
			},

			filename: "go1.20.3#macos|arm32@tar.gz",
		},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("format %s", p.format), func(t *testing.T) {
			filename, err := Format(p.format, p.version)
			if err != p.err {
				t.Errorf("expected error %v, got %v", p.err, err)
			}
			if filename != p.filename {
				t.Errorf("expected filename %v, got %v", p.filename, filename)
			}
		})
	}
}
