package env

import "path/filepath"

type Variables struct {
	// ConfigFile path where variables below can be loaded from.
	ConfigFile string

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
