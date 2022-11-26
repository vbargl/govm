package resources

import (
	_ "embed"
)

var (
	// GoDevDlContent represents content of https://go.dev/dl downloaded at 2023-08-03
	//go:embed view-source_https___go.dev_dl_.html
	GoDevDlContent string
)
