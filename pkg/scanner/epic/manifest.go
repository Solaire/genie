package epic

import (
	"encoding/json"
	"os"
)

type EpicManifest struct {
	Name        string `json:"MainGameAppName"`  // Game name. Not sure how this is different from DisplayName.
	DisplayName string `json:"DisplayName"`      // Game name. Not sure how this is different from Name.
	InstallDir  string `json:"InstallLocation"`  // Root dir of the game.
	LaunchDir   string `json:"LaunchExecutable"` // Relative path to game exe.
}

func parseManifest(path string) (*EpicManifest, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var manifest EpicManifest
	if err := json.Unmarshal(raw, &manifest); err != nil {
		return nil, err

	}

	return &manifest, nil
}
