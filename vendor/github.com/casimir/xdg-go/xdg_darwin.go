package xdg

import "os"

func DataHome() string {
	return os.Getenv("HOME") + "/Library"
}

func ConfigHome() string {
	return os.Getenv("HOME") + "/Library/Preferences"
}

func CacheHome() string {
	return os.Getenv("HOME") + "/Library/Caches"
}

func DataDirs() []string {
	return []string{"/Library"}
}

func ConfigDirs() []string {
	return []string{"/Library/Preferences", "/Library/Application Support"}
}
