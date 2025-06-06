//go:build linux
// +build linux

package battlenet

import (
	"errors"

	"github.com/solaire/genie/internal/db"
)

func findBattleNetLauncher() (string, error) {
	// NOTE: Battle.net does not support Linux
	return "", errors.New("Battle.net not supported on Linux")
}

func findBattleNetCache(client_root string) (string, error) {
	// NOTE: Battle.net does not support Linux
	return "", errors.New("Battle.net not supported on Linux")
}

func getCacheFiles(cache string) ([]string, error) {
	// NOTE: Battle.net does not support Linux
	return nil, errors.New("Battle.net not supported on Linux")
}

func buildGameData(data *UninstallData) (*db.Game, error) {
	// NOTE: Battle.net does not support Linux
	return nil, errors.New("Battle.net not supported on Linux")
}
