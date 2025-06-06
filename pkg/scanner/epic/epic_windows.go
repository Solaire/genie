//go:build windows
// +build windows

package epic

import (
	"io/fs"
	"path/filepath"

	"github.com/solaire/genie/internal/utils"

	"golang.org/x/sys/windows/registry"
)

func findEpicGamesLauncher() (string, string, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Epic Games\EOS`, registry.QUERY_VALUE)
	if err != nil {
		return "", "", err
	}
	defer k.Close()

	base_bin, err := utils.RegistryGetStringPath(k, "ModSdkCommand")
	if err != nil {
		return "", "", err
	}

	manifest_dir, err := utils.RegistryGetStringPath(k, "ModSdkMetadataDir")
	if err != nil {
		return "", "", err
	}

	return base_bin, manifest_dir, nil
}

func findEpicGamesManifests(path string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".item" {
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
