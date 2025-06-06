//go:build windows
// +build windows

package battlenet

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/solaire/genie/internal/utils"
	"github.com/solaire/genie/pkg/models"

	"golang.org/x/sys/windows/registry"
)

func findBattleNetLauncher() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall\Battle.net`, registry.QUERY_VALUE)
	if err != nil {
		return "", nil
	}
	defer k.Close()

	path, err := utils.RegistryGetStringPath(k, `InstallLocation`)
	if err != nil {
		return "", nil
	}
	return path, nil
}

func findBattleNetCache(client_root string) (string, error) {
	const path = `C:\ProgramData\Battle.net\Agent\data\cache`
	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}

func getCacheFiles(cache string) ([]string, error) {
	var cache_files []string

	err := filepath.WalkDir(cache, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			if filepath.Ext(d.Name()) == "" {
				cache_files = append(cache_files, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return cache_files, nil
}

func buildGameData(data *UninstallData) (*models.Game, error) {
	// Find out the right registry key
	idx := 0
	for i, reg := range data.RegistryData {
		if reg.KeyType == "HKEY_LOCAL_MACHINE" {
			idx = i
			break
		}
	}

	if len(data.RegistryData[idx].Subkeys) == 0 {
		return nil, fmt.Errorf("no subkeys found in the registry data")
	}

	kp := filepath.Join(data.RegistryData[idx].Root, data.RegistryData[idx].Subkeys[0])
	if data.RegistryData[idx].Flags == "WOW_BOTH" {
		kp = strings.Replace(strings.ToLower(kp), "software\\", "software\\wow6432node\\", 1)
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, kp, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	game := &models.Game{
		Name:       "",
		Platform:   "battle.net",
		Path:       "",
		Executable: "",
		LaunchCmd:  "",
	}

	// DisplayName
	if game.Name, _, err = k.GetStringValue("DisplayName"); err != nil {
		return nil, err
	}

	// InstallLocation
	if game.Path, err = utils.RegistryGetStringPath(k, "InstallLocation"); err != nil {
		return nil, err
	}

	// DisplayIcon as executable
	if game.Executable, err = utils.RegistryGetStringPath(k, "DisplayIcon"); err != nil {
		return nil, err
	}

	return game, nil
}
