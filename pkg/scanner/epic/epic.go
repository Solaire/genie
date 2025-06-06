package epic

import "github.com/solaire/genie/pkg/models"

type Scanner struct {
	Launcher     string
	ManifestPath string
}

func (s *Scanner) Name() string {
	return "epic"
}

func (s *Scanner) Detect() bool {
	bin, manifest_path, err := findEpicGamesLauncher()
	if err != nil {
		return false
	}
	s.Launcher = bin
	s.ManifestPath = manifest_path
	return true
}

func (s *Scanner) ScanGames() ([]models.Game, error) {
	manifests, err := findEpicGamesManifests(s.ManifestPath)
	if err != nil {
		return nil, err
	}

	var games []models.Game

	for _, item := range manifests {
		parsed, err := parseManifest(item)
		if err != nil {
			return nil, err
		}

		game := models.Game{
			Name:       parsed.DisplayName,
			Platform:   "epic",
			Path:       parsed.InstallDir,
			Executable: parsed.LaunchDir,
			LaunchCmd:  "",
		}
		games = append(games, game)
	}

	return games, nil
}
