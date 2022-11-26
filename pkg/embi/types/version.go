package types

// Version is verified go version in format go1.2.3 including beta and rc types
type Version string

func (v Version) String() string {
	return string(v)
}

// QualifiedVersion is version with targeted platform (goos and goarch)
type QualifiedVersion struct {
	Version Version // required
	Goos    string  // if empty build.Default.GOOS will be used
	Goarch  string  // if empty build.Default.GOARCH will be used
}

// TODO: move these two below somewhere else

type VersionSet struct {
	Version Version
	Files   []VersionFile
}

type VersionFile struct {
	Version  Version
	Url      string
	Goarch   string
	Goos     string
	Checksum string
}
