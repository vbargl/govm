package env

import (
	"log"
	"os"
	"path/filepath"

	"barglvojtech.net/govm/pkg/internal/config"

	"barglvojtech.net/x/pkg/stringutil"
)

// Init function will populate Variables
func Init(configDir string) Variables {
	configDir = stringutil.FirstNonEmpty(configDir, DefaultConfigPath)
	configFile := filepath.Join(configDir, "config.toml")

	vars := Default()
	overrideWithConfig(&vars, configFile)
	return vars
}

func overrideWithConfig(vars *Variables, configFile string) {
	if _, err := os.Stat(configFile); err != nil {
		return
	}

	cfg, err := config.Parse(configFile)
	if err != nil {
		log.Printf("config file could not be loaded %s: %v", configFile, err)
		return
	}

	vars.ConfigFile = configFile // always override
	overrideIfPresent(&vars.BinDir, cfg.Core.BinDir)
	overrideIfPresent(&vars.RuntimeDir, cfg.Core.RuntimeDir)
}

func overrideIfPresent[T comparable](target *T, value T) {
	var zero T

	if value != zero {
		*target = value
	}
}
