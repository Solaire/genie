package battlenet

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/solaire/genie/internal/utils"
)

type UninstallData struct {
	RegistryData []RegistryEntry `json:"delete_registry_key_list"`
}

type RegistryEntry struct {
	Flags   string   `json:"flags"`
	KeyType string   `json:"key_type"`
	Root    string   `json:"root"`
	Subkeys []string `json:"subkeys"`
}

func parseCacheFile(file string) (*UninstallData, error) {
	raw, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var full_json map[string]any
	if err := json.Unmarshal(raw, &full_json); err != nil {
		return nil, err
	}

	config, err := utils.JsonExtractInner(full_json, "platform", "win", "config")
	if err != nil {
		return nil, err
	}

	uninstall, ok := config["uninstall"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("uninstall not found in cache file")
	}

	uninstall_bytes, err := json.Marshal(uninstall)
	if err != nil {
		return nil, err
	}

	registry_data, err := extractRegistryData(uninstall_bytes)
	if err != nil {
		return nil, err
	}

	if len(registry_data.RegistryData) == 0 {
		return nil, fmt.Errorf("delete_registry_key_list not found in cache file")
	}

	return registry_data, nil
}

func extractRegistryData(data []byte) (*UninstallData, error) {
	// Decode array of objects with dynamic keys
	var items []map[string]json.RawMessage
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	var uninstallData UninstallData
	for _, item := range items {
		if val, ok := item["delete_registry_key_list"]; ok {
			var entry RegistryEntry
			if err := json.Unmarshal(val, &entry); err != nil {
				return nil, err
			}
			uninstallData.RegistryData = append(uninstallData.RegistryData, entry)
		}
	}

	return &uninstallData, nil
}
