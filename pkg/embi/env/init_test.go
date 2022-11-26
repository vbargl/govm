package env

import (
	"testing"

	"barglvojtech.net/govm/pkg/internal/versionutil"
)

func TestInit(t *testing.T) {
	config := Init("")

	assert(t, config.BinDir == DefaultBinPath, "BinDir")
	assert(t, config.RuntimeDir == DefaultRuntimePath, "RuntimeDir")
	assert(t, config.ConfigFile == DefaultConfigPath, "RuntimeDir")
	assert(t, config.VersionUrl == versionutil.DefaultVersionUrl, "RuntimeDir")
}

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Error(msg)
	}
}
