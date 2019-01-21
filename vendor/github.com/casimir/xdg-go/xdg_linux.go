package xdg

import "os"

func DataHome() string {
	return os.Getenv("HOME") + "/.local/share"
}

func ConfigHome() string {
	return os.Getenv("HOME") + "/.config"
}

func CacheHome() string {
	return os.Getenv("HOME") + "/.cache"
}

func DataDirs() []string {
	// The specification gives a  value with trailing slashes but only
	// for this value. Seems odd enough to take the liberty of removing them.
	return []string{"/usr/local/share", "/usr/share"}
}

func ConfigDirs() []string {
	return []string{"/etc/xdg"}
}
