//go:build linux
// +build linux

package gog

import (
	"errors"

	"github.com/solaire/genie/internal/db"
)

func findGogPath() (string, error) {
	// NOTE: Gog galaxy does not support Linux
	return "", errors.New("GOG not supported on Linux")
}

func findGogGameInfoFiles() ([]string, error) {
	// NOTE: Gog galaxy does not support Linux
	return nil, errors.New("GOG not supported on Linux")
}

func findInstalledGames() ([]db.Game, error) {
	// NOTE: Gog galaxy does not support Linux
	return nil, errors.New("GOG not supported on Linux")
}
