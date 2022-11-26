package env

import "path/filepath"

type Variables struct {
	// ConfigFile path where variables below can be loaded from.
	ConfigFile string

	// VersionUrl which contains variables for version goos goarch and ext encapsulated in `${}`.
	// If no url is specified, default one will be used: https://go.dev/dl/${version}.${goos}-${goarch}.${ext}.
	VersionUrl string

	// RuntimeDir where sdk will be extracted into.
	RuntimeDir string

	// BinDir where to put symbolic link of go binary.
	BinDir string
}

func (vars *Variables) CacheDir() string {
	return filepath.Join(vars.RuntimeDir, "cache")
}

func (vars *Variables) VersionsDir() string {
	return filepath.Join(vars.RuntimeDir, "versions")
}

func (vars *Variables) GoExecutable() string {
	return filepath.Join(vars.BinDir, "go")
}
