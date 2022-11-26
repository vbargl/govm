// go:build linux

package env

import (
	"os"
	"path/filepath"

	"barglvojtech.net/x/pkg/errutil"
)

var (
	DefaultRuntimePath string
	DefaultConfigPath  string
	DefaultBinPath     string
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	errutil.LogFatalIfErr(err, errutil.PrefixWith("could not obtain user home directory"))
	DefaultRuntimePath = filepath.Join(userHomeDir, ".local", "share", "govm")
	DefaultBinPath = filepath.Join(userHomeDir, ".local", "bin")

	userConfigDir, err := os.UserConfigDir()
	errutil.LogFatalIfErr(err, errutil.PrefixWith("could not obtain user config directory"))
	DefaultConfigPath = filepath.Join(userConfigDir, "govm")
}
