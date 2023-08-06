package env

import (
	"testing"
)

func TestInit(t *testing.T) {
	config := Init("")

	assert(t, config.BinDir == DefaultBinPath, "BinDir")
	assert(t, config.RuntimeDir == DefaultRuntimePath, "RuntimeDir")
	assert(t, config.ConfigFile == DefaultConfigPath, "ConfigDir")
}

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Error(msg)
	}
}
