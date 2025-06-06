//go:build windows
// +build windows

package gog

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/solaire/genie/internal/utils"
	"github.com/solaire/genie/pkg/models"

	"golang.org/x/sys/windows/registry"
)

func findGogPath() (string, error) {
	config_json := `C:\ProgramData\GOG.com\Galaxy\config.json`
	_, err := os.Stat(config_json)
	if err != nil {
		return "", err
	}

	raw, err := os.ReadFile(config_json)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		return "", err
	}

	if path, ok := data["libraryPath"].(string); ok {
		return path, nil
	}

	return "", errors.New("GOG not found")
}

func findGogGameInfoFiles() ([]string, error) {
	root := `C:\ProgramData\GOG.com\supportInstaller`
	if _, err := os.Stat(root); err != nil {
		return nil, err
	}

	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".info" {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func findInstalledGames() ([]models.Game, error) {
	const root = `SOFTWARE\WOW6432Node\GOG.com\Games`

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, root, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	subkeys, err := k.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}

	var games []models.Game

	for _, key := range subkeys {
		game_key, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Join(root, key), registry.QUERY_VALUE)
		if err != nil {
			return nil, err
		}

		// Name
		name, _, err := game_key.GetStringValue("gameName")
		if err != nil {
			return nil, err
		}

		// Install location
		path, err := utils.RegistryGetStringPath(game_key, "path")
		if err != nil {
			return nil, err
		}

		// executable
		executable, err := utils.RegistryGetStringPath(game_key, "launchCommand")
		if err != nil {
			return nil, err
		}

		game := models.Game{
			Name:       name,
			Platform:   "gog",
			Path:       path,
			Executable: executable,
			LaunchCmd:  "",
		}

		games = append(games, game)
	}

	return games, nil
}
