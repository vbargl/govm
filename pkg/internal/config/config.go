package config

// Config represents
type Config struct {
	Core Core
}

type Core struct {
	// RuntimeDir where sdk will be extracted into.
	RuntimeDir string

	// BinDir where to put symbolic link of go binary.
	BinDir string
}
