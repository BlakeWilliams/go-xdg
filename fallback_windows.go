package xdg

import (
	"fmt"
	"os"
	"path"
)

var fallbackEnvMap = map[xdg]string{
	xdgConfig: "LOCALAPPDATA",
	xdgData:   "LOCALAPPDATA",
	xdgCache:  "TEMP",
	xdgState:  "LOCALAPPDATA",
}

var fallbackMap = map[xdg][]string{
	xdgConfig: {"AppData", "Local"},
	xdgData:   {"AppData", "Local"},
	xdgCache:  {"AppData", "Local", "Temp"},
	xdgState:  {"AppData", "Local"},
}

func fallback(target string) (string, error) {
	if env := os.Getenv(fallbackEnvMap[target]); env != "" {
		return env, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user's home directory: %w", err)
	}

	return path.Join(home, path.Join(fallbackMap[target]...)), nil
}
