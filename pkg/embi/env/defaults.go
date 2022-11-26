package env

import "barglvojtech.net/govm/pkg/internal/versionutil"

// Default returns default values configured by embi package.
// Also see defaults_${GOOS}.go for aditional defaults for specific GOOS.
func Default() Variables {
	return Variables{
		ConfigFile: DefaultConfigPath,
		VersionUrl: versionutil.DefaultVersionUrl,
		RuntimeDir: DefaultRuntimePath,
		BinDir:     DefaultBinPath,
	}
}
