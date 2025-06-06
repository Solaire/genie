//go:build linux
// +build linux

package epic

import "errors"

func findEpicGamesLauncher() (string, string, error) {
	// NOTE: Epic client does not support Linux
	return "", "", errors.New("Epic client not supported on Linux")
}

func findEpicGamesManifests(path string) ([]string, error) {
	// NOTE: Epic client does not support Linux
	return nil, errors.New("Epic client not supported on Linux")
}
