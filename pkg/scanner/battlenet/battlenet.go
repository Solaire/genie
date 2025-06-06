package battlenet

import (
	"github.com/solaire/genie/pkg/logger"
	"github.com/solaire/genie/pkg/models"
)

type Scanner struct {
	InstallDir string
	CacheDir   string
}

func (s *Scanner) Name() string {
	return "battle.net"
}

func (s *Scanner) Detect() bool {
	path, err := findBattleNetLauncher()
	if err != nil {
		return false
	}
	s.InstallDir = path

	cache_root, err := findBattleNetCache(s.InstallDir)
	if err != nil {
		return false
	}
	s.CacheDir = cache_root

	return true
}

func (s *Scanner) ScanGames() ([]models.Game, error) {

	cache_files, err := getCacheFiles(s.CacheDir)
	if err != nil {
		return nil, err
	}

	var games []models.Game

	for _, cache_file := range cache_files {
		parsed, err := parseCacheFile(cache_file)
		if err != nil {
			logger.Printf("ERROR parsing cache file: %v\n", err)
			continue
		}

		game, err := buildGameData(parsed)
		if err != nil {
			logger.Printf("ERROR building game data: %v\n", err)
			continue
		}

		games = append(games, *game)
	}

	return games, nil
}
