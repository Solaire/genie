package ea

import "github.com/solaire/genie/pkg/models"

// Used to be called Origin...
// ... now Scanner App
type Scanner struct {
	InstallDir string
}

func (s *Scanner) Name() string {
	return "ea"
}

func (s *Scanner) Detect() bool {
	path, err := findEaLauncher()
	if err != nil {
		return false
	}
	s.InstallDir = path
	return true
}

func (s *Scanner) ScanGames() ([]models.Game, error) {
	manifest, err := findEaManifest()
	if err != nil {
		return nil, err
	}

	manifest_data, err := decryptManifest(manifest)
	if err != nil {
		return nil, err
	}

	var games []models.Game

	for _, info := range manifest_data.InstallInfos {
		game, err := getEaGameData(info)
		if err != nil {
			return nil, err
		}
		games = append(games, *game)
	}

	return games, nil
}
