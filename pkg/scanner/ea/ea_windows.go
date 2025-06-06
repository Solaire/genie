//go:build windows
// +build windows

package ea

import (
	"os"
	"strings"

	"github.com/solaire/genie/internal/utils"
	"github.com/solaire/genie/pkg/models"

	"golang.org/x/sys/windows/registry"
)

// Find the launcher binary for EA App.
// Seems like the app is always forced to download in C:/Program Files/...
func findEaLauncher() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `software\Electronic Arts\EA Desktop`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	// There are 2 possible binaries:
	// - EALauncher.exe (LauncherAppPath)
	// - EADesktop.exe (DesktopAppPath)
	// Both seem to do the same thing...
	launcher_dir, err := utils.RegistryGetStringPath(k, "InstallLocation")
	if err != nil {
		return "", err
	}

	return launcher_dir, nil
}

// Find the location of the manifest file
func findEaManifest() (string, error) {
	const manifest_path = `C:\ProgramData\EA Desktop\530c11479fe252fc5aabc24935b9776d4900eb3ba58fdc271e0d6229413ad40e\IS` // Same for each user
	_, err := os.Stat(manifest_path)
	if err != nil {
		return "", err
	}
	return manifest_path, nil
}

func getEaGameData(info InstallInfo) (*models.Game, error) {
	reg_key := info.InstallReg
	idx := strings.Index(reg_key, "]")
	if idx > 0 {
		reg_key = reg_key[:idx]
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, reg_key, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	name, err := utils.RegistryGetStringPath(k, "displayname")
	if err != nil {
		return nil, err
	}

	return &models.Game{
		Name:     name,
		Platform: "ea",
		Path:     info.InstallPath,
	}, nil
}
