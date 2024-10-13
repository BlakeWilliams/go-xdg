package xdg

import (
	"os"
)

type xdg = string

const (
	xdgConfig xdg = "XDG_CONFIG_HOME"
	xdgData   xdg = "XDG_DATA_HOME"
	xdgCache  xdg = "XDG_CACHE_HOME"
	xdgState  xdg = "XDG_STATE_HOME"
)

// ConfigHome returns the path to the user's configuration directory.
// It tries the following environment variables in order:
//   - $XDG_CONFIG_HOME
//   - $HOME/.config
func ConfigHome() (string, error) {
	if path := os.Getenv(xdgConfig); path != "" {
		return path, nil
	}

	return fallback(xdgConfig)
}

// DataHome returns the path to the user's data directory.
// It tries the following environment variables in order:
//   - $XDG_DATA_HOME
//   - $HOME/.local/share
func DataHome() (string, error) {
	if path := os.Getenv(xdgData); path != "" {
		return path, nil
	}

	return fallback(xdgData)
}

// CacheHome returns the path to the user's cache directory.
// It tries the following environment variables in order:
//   - $XDG_CACHE_HOME
//   - $HOME/.cache
func CacheHome() (string, error) {
	if path := os.Getenv(xdgCache); path != "" {
		return path, nil
	}

	return fallback(xdgCache)
}

// StateHome returns the path to the user's cache directory.
// It tries the following environment variables in order:
//   - $XDG_STATE_HOME
//   - $HOME/.cache
func StateHome() (string, error) {
	if path := os.Getenv(xdgState); path != "" {
		return path, nil
	}

	return fallback(xdgState)
}
