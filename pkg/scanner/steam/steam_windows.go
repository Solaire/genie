//go:build windows
// +build windows

package steam

import (
	"os"

	"golang.org/x/sys/windows/registry"
)

func findSteamPath() (string, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Valve\Steam`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	path, _, err := k.GetStringValue("SteamPath")
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}
