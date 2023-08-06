package env

// Default returns default values configured by embi package.
// Also see defaults_${GOOS}.go for aditional defaults for specific GOOS.
func Default() Variables {
	return Variables{
		ConfigFile: DefaultConfigPath,
		RuntimeDir: DefaultRuntimePath,
		BinDir:     DefaultBinPath,
	}
}
