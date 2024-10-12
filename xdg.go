package xdg

import (
	"fmt"
	"os"
	"path"
)

// ConfigHome returns the path to the user's configuration directory.
// It tries the following environment variables in order:
//   - $XDG_CONFIG_HOME
//   - $HOME/.config
func ConfigHome() (string, error) {
	if path := os.Getenv("XDG_CONFIG_HOME"); path != "" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user's home directory: %w", err)
	}

	return path.Join(home, ".config"), nil
}

// DataHome returns the path to the user's data directory.
// It tries the following environment variables in order:
//   - $XDG_DATA_HOME
//   - $HOME/.local/share
func DataHome() (string, error) {
	if path := os.Getenv("XDG_DATA_HOME"); path != "" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user's home directory: %w", err)
	}

	return path.Join(home, ".local/share"), nil
}

// CacheHome returns the path to the user's cache directory.
// It tries the following environment variables in order:
//   - $XDG_CACHE_HOME
//   - $HOME/.cache
func CacheHome() (string, error) {
	if path := os.Getenv("XDG_CACHE_HOME"); path != "" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user's home directory: %w", err)
	}

	return path.Join(home, ".cache"), nil
}

// StateHome returns the path to the user's cache directory.
// It tries the following environment variables in order:
//   - $XDG_STATE_HOME
//   - $HOME/.cache
func StateHome() (string, error) {
	if path := os.Getenv("XDG_STATE_HOME"); path != "" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user's home directory: %w", err)
	}

	return path.Join(home, ".local/state"), nil
}
