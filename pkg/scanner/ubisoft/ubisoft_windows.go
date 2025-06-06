//go:build windows
// +build windows

package ubisoft

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/solaire/genie/internal/utils"

	"golang.org/x/sys/windows/registry"
)

// Find the Installation directory for Ubisoft client
func findUbisoftPath() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `software\WOW6432Node\ubisoft\Launcher`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	install_dir, _, err := k.GetStringValue("InstallDir")
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(install_dir); err != nil {
		return "", err
	}

	return install_dir, nil
}

// Find the IDs and installation directories of Ubisoft games.
func findUbisoftInstalls() (map[string]string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `software\WOW6432Node\ubisoft\Launcher\Installs`, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	subkeys, err := k.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}

	installs := make(map[string]string)

	for _, key := range subkeys {
		k2, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(`software\WOW6432Node\ubisoft\Launcher\Installs`, key), registry.QUERY_VALUE)
		if err != nil {
			return nil, err
		}

		install_dir, _, err := k2.GetStringValue("InstallDir")
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(install_dir); err != nil {
			return nil, err
		}

		installs[key] = install_dir
	}

	return installs, nil
}

// Retrieve the game titles from the registry
func findUbisoftGameData(installs map[string]string) (map[string]string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	subkeys, err := k.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}
	subkeys = utils.FilterArray(subkeys, func(s string) bool {
		return strings.Contains(s, "Uplay Install")
	})

	data := make(map[string]string)

	for _, entry := range subkeys {
		k2, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(`software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`, entry), registry.QUERY_VALUE)
		if err != nil {
			return nil, err
		}

		id := strings.TrimSpace(strings.TrimPrefix(entry, "Uplay Install"))
		name, _, err := k2.GetStringValue("DisplayName")
		if err != nil {
			return nil, err
		}

		data[id] = name
	}

	return data, nil
}

/*
func findUbisoftInstalls() (map[string][2]string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	subkeys, err := k.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}
	subkeys = utils.Filter(subkeys, func(s string) bool {
		return strings.Contains(s, "Uplay Install")
	})

	installs := make(map[string][2]string)
	for _, entry := range subkeys {
		k2, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(`software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`, entry), registry.QUERY_VALUE)
		if err != nil {
			return nil, err
		}

		id := strings.TrimSpace(strings.TrimPrefix(entry, "Uplay Install"))

		name, _, err := k2.GetStringValue("DisplayName")
		if err != nil {
			return nil, err
		}

		dir, _, err := k2.GetStringValue("InstallLocation")
		if err != nil {
			return nil, err
		}
		if _, err := os.Stat(dir); err != nil {
			return nil, err
		}

		installs[id] = [2]string{
			name,
			dir,
		}

		k2.Close()
	}

	return installs, nil
}
*/

/*
func findUbisoftInstalls() (map[string]string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `software\WOW6432Node\ubisoft\Launcher\Installs`, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	subkeys, err := k.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}

	installs := make(map[string]string)
	for _, key := range subkeys {
		k2, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(`software\WOW6432Node\ubisoft\Launcher\Installs`, key), registry.QUERY_VALUE)
		if err != nil {
			return nil, err
		}

		path, _, err := k2.GetStringValue("InstallDir")
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(path); err != nil {
			return nil, nil
		}

		installs[key] = path
		k2.Close()
	}

	return installs, nil
}
*/
