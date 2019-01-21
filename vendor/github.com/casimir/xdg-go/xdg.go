package xdg

import (
	"os"
	"path/filepath"
	"strings"
)

// App is an aggregation of XDG information for a named application. Useful to
// propagate app configuration.
type App struct {
	Name string
}

func (a App) path(envVar string, defaultFn func() string, file string) string {
	base := os.Getenv(envVar)
	if base == "" {
		base = defaultFn()
	}
	return filepath.Join(base, a.Name, file)
}

// DataPath determines the full path of a data file.
func (a App) DataPath(file string) string {
	return a.path("XDG_DATA_HOME", DataHome, file)
}

// ConfigPath determines the full path of a data file.
func (a App) ConfigPath(file string) string {
	return a.path("XDG_CONFIG_HOME", ConfigHome, file)
}

// CachePath determines the full path of a cached file.
func (a App) CachePath(file string) string {
	return a.path("XDG_CACHE_HOME", CacheHome, file)
}

func (a App) multiPaths(envVar string, defaultFn func() []string, file string) []string {
	var bases []string
	env := os.Getenv(envVar)
	if env != "" {
		bases = strings.Split(env, ":")
	} else {
		bases = defaultFn()
	}
	var dirs []string
	for _, it := range bases {
		dirs = append(dirs, filepath.Join(it, a.Name, file))
	}
	return dirs
}

// SystemDataPaths determines system-wide possible paths for a data file.
func (a App) SystemDataPaths(file string) []string {
	return a.multiPaths("XDG_DATA_DIRS", DataDirs, file)
}

// SystemConfigPaths determines system-wide possible paths for a config file.
func (a App) SystemConfigPaths(file string) []string {
	return a.multiPaths("XDG_CONFIG_DIRS", ConfigDirs, file)
}

var defaultApp App

// SetName for the default application. Used for package-wide functions.
func SetName(name string) {
	defaultApp.Name = name
}

// DataPath determines the full path of a data file.
func DataPath(file string) string {
	return defaultApp.DataPath(file)
}

// ConfigPath determines the full path of a data file.
func ConfigPath(file string) string {
	return defaultApp.ConfigPath(file)
}

// CachePath determines the full path of a cached file.
func CachePath(file string) string {
	return defaultApp.CachePath(file)
}

// SystemDataPaths determines system-wide possible paths for a data file.
func SystemDataPaths(file string) []string {
	return defaultApp.SystemDataPaths(file)
}

// SystemConfigPaths determines system-wide possible paths for a config file.
func SystemConfigPaths(file string) []string {
	return defaultApp.SystemConfigPaths(file)
}
