//go:build linux
// +build linux

package ubisoft

import "errors"

func findUbisoftPath() (string, error) {
	// NOTE: Ubisoft client does not support Linux
	return "", errors.New("Ubisoft client not supported on Linux")
}

func findUbisoftInstalls() (map[string]string, error) {
	// NOTE: Ubisoft client does not support Linux
	return nil, errors.New("Ubisoft client not supported on Linux")
}

func findUbisoftGameData(installs map[string]string) (map[string]string, error) {
	// NOTE: Ubisoft client does not support Linux
	return nil, errors.New("Ubisoft client not supported on Linux")
}
