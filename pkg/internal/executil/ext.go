package executil

func Ext(goos string) string {
	switch goos {
	case "windows":
		return ".exe"
	default:
		return ""
	}
}
