//go:build linux
// +build linux

package steam

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

func findSteamPath() (string, error) {
	usr, _ := user.Current()
	paths := []string{
		filepath.Join(usr.HomeDir, ".steam", "steam"),
		filepath.Join(usr.HomeDir, ".steam", "share", "Steam"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	return "", errors.New("Steam not found")
}
