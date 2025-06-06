//go:build linux
// +build linux

package ea

import (
	"errors"

	"github.com/solaire/genie/internal/db"
)

func findEaLauncher() (string, error) {
	// NOTE: EA App does not support Linux
	return "", errors.New("EA App not supported on Linux")
}

func findEaManifest() (string, error) {
	// NOTE: EA App does not support Linux
	return "", errors.New("EA App not supported on Linux")
}

func getEaGameData(info InstallInfo) (*db.Game, error) {
	// NOTE: EA App does not support Linux
	return nil, errors.New("EA App not supported on Linux")
}
