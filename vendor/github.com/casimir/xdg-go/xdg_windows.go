package xdg

import "os"

func DataHome() string {
	return os.Getenv("APPDATA")
}

func ConfigHome() string {
	return os.Getenv("APPDATA")
}

func CacheHome() string {
	return os.Getenv("TEMP")
}

func DataDirs() []string {
	return []string{os.Getenv("APPDATA"), os.Getenv("LOCALAPPDATA")}
}

func ConfigDirs() []string {
	return []string{os.Getenv("APPDATA"), os.Getenv("LOCALAPPDATA")}
}
