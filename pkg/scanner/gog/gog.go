package gog

import "github.com/solaire/genie/pkg/models"

type Scanner struct {
	InfoFiles []string
	BasePath  string
}

func (s *Scanner) Name() string {
	return "gog"
}

func (s *Scanner) Detect() bool {
	base, err := findGogPath()
	if err != nil {
		return false
	}
	s.BasePath = base

	info_files, err := findGogGameInfoFiles()
	if err != nil {
		return false
	}

	s.InfoFiles = info_files
	return true
}

func (s *Scanner) ScanGames() ([]models.Game, error) {
	games, err := findInstalledGames()
	if err != nil {
		return nil, err
	}
	return games, nil
}

/*
func (s *GogScanner) ScanGames() ([]db.Game, error) {
	var games []db.Game
	info_map := make(map[string]*InfoFile)

	// Unmarshal each info file and
	// make a binary-InfoFile map
	for _, info_file := range s.InfoFiles {
		info, err := parseInfo(info_file)
		if err != nil {
			return nil, err
		}

		for _, task := range info.PlayTasks {
			if task.IsPrimary {
				info_map[task.Path] = info
				break
			}
		}
	}

	// The metadata doesn't store the game's directory.
	// Knowing the root directory and the name of the binary,
	// we can find the game's directory.

	entries, err := os.ReadDir(s.BasePath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		subdir := filepath.Join(s.BasePath, entry.Name())
		found := ""

		for binary, info := range info_map {
			bin_path := filepath.Join(subdir, binary)
			if _, err := os.Stat(bin_path); err == nil {
				game := db.Game{
					Name:          info.Name,
					Platform:      "gog",
					Path:          subdir,
					Executable:    binary,
					LaunchOptions: "",
				}
				games = append(games, game)
				found = binary
				break
			}
		}

		if found != "" {
			delete(info_map, found)
		}
	}
	return games, nil
}
*/
