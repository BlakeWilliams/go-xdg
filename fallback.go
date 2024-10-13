//go:build !windows

package xdg

import (
	"fmt"
	"os"
	"path"
)

var fallbackMap = map[xdg][]string{
	xdgConfig: {".config"},
	xdgData:   {".local", "share"},
	xdgCache:  {".cache"},
	xdgState:  {".local", "state"},
}

func fallback(target string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user's home directory: %w", err)
	}

	return path.Join(home, path.Join(fallbackMap[target]...)), nil
}
