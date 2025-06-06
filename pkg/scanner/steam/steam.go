package steam

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/solaire/genie/pkg/models"
)

type Scanner struct {
	LibraryFolders []string
	BasePath       string
}

func (s *Scanner) Name() string {
	return "steam"
}

func (s *Scanner) Detect() bool {
	base, err := findSteamPath()
	if err != nil {
		return false
	}

	s.BasePath = base
	vdf_path := filepath.Join(base, "steamapps", "libraryfolders.vdf")
	libraries, err := parseLibraryFoldersFile(vdf_path)

	//libraries, err := getSteamLibraries(base)
	if err != nil {
		return false
	}

	s.LibraryFolders = libraries
	return true
}

func (s *Scanner) ScanGames() ([]models.Game, error) {
	var games []models.Game

	for _, lib := range s.LibraryFolders {
		manifests, _ := filepath.Glob(filepath.Join(lib, "steamapps", "appmanifest_*.acf"))
		for _, mf := range manifests {
			game, err := parseAppManifestFile(mf, lib)
			if err == nil {
				games = append(games, *game)
			}
		}
	}

	return games, nil
}

func parseLibraryFoldersFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := parseVdf(f)
	if err != nil {
		return nil, err
	}

	var paths []string
	if libraries, ok := data["libraryfolders"].(VdfNode); ok {
		for _, v := range libraries {
			if entry, ok := v.(VdfNode); ok {
				if path, ok := entry["path"].(string); ok {
					paths = append(paths, path)
				}
			}
		}
	}

	return paths, nil
}

func parseAppManifestFile(manifestPath, libraryFolder string) (*models.Game, error) {
	f, err := os.Open(manifestPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := parseVdf(f)
	if err != nil {
		return nil, err
	}

	root, ok := data["AppState"].(VdfNode)
	if !ok {
		return nil, fmt.Errorf("invalid manifest format: missing AppState")
	}

	getStr := func(key string) (string, error) {
		val, ok := root[key]
		if !ok {
			return "", fmt.Errorf("missing key %q in %s", key, manifestPath)
		}

		s, ok := val.(string)
		if !ok {
			return "", fmt.Errorf("key %q is not a string", key)
		}

		return s, nil
	}

	app_id_str, err := getStr("appid")
	if err != nil {
		return nil, err
	}

	name, err := getStr("name")
	if err != nil {
		return nil, err
	}

	install_dir, err := getStr("installdir")
	if err != nil {
		return nil, err
	}

	game_dir := filepath.Join(libraryFolder, "common", install_dir)
	return &models.Game{
		Name:       name,
		Platform:   "steam",
		Path:       game_dir,
		Executable: "",
		LaunchCmd:  fmt.Sprintf("steam://rungameid/%s", app_id_str), // Steam's Game ID for launching
	}, nil
}
