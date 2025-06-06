package gog

import (
	"encoding/json"
	"os"
)

type InfoFile struct {
	ClientID   string     `json:"clientId"`
	GameID     string     `json:"gameId"`
	Language   string     `json:"language"`
	Languages  []string   `json:"languages"`
	Name       string     `json:"name"`
	PlayTasks  []PlayTask `json:"playTasks"`
	RootGameID string     `json:"rootGameId"`
	Version    int        `json:"version"`
}

type PlayTask struct {
	Category  string   `json:"category"`
	IsPrimary bool     `json:"isPrimary,omitempty"`
	Languages []string `json:"languages"`
	Link      string   `json:"link,omitempty"`
	Name      string   `json:"name"`
	Path      string   `json:"path,omitempty"`
	Type      string   `json:"type"`
}

func parseInfo(path string) (*InfoFile, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var info InfoFile
	if err := json.Unmarshal(raw, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
