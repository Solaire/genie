package utils

import (
	"os"

	"golang.org/x/sys/windows/registry"
)

// Retrieve string value from the registry, assuming it's a path to a directory
// or a file. Perform an additional 'os.Stat' to ensure it exists.
func RegistryGetStringPath(k registry.Key, value string) (string, error) {
	path, _, err := k.GetStringValue(value)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}
