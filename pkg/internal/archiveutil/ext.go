package archiveutil

const (
	fileExtensionWindows = "zip"
	fileExtensionUnix    = "tar.gz"
)

func GetExtension(goos string) string {
	switch goos {
	case "windows":
		return fileExtensionWindows
	default:
		return fileExtensionUnix
	}
}
