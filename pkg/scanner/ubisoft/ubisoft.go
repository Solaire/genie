package ubisoft

import (
	"fmt"

	"github.com/solaire/genie/internal/utils"
	"github.com/solaire/genie/pkg/models"
)

type Scanner struct {
	BasePath string
}

func (s *Scanner) Name() string {
	return "ubisoft"
}

func (s *Scanner) Detect() bool {
	base, err := findUbisoftPath()
	if err != nil {
		return false
	}
	s.BasePath = base
	return true
}

func (s *Scanner) ScanGames() ([]models.Game, error) {
	installs, err := findUbisoftInstalls()
	if err != nil {
		return nil, err
	}

	// At this point we have the Ubisoft base directory, as well
	//   as the ID and directory of each installed game.
	// Now we need to find the name of the game and binary.
	// Try to get the cached data from "configurations" file, and
	//   fallback to registry if needed.
	var games []models.Game

	if cache, err := parseConfigurationsFile(s.BasePath); err == nil {
		// Cache can be pretty big. Remove uninstalled games before parsing the YAML
		cache = utils.FilterArray(cache, func(s UplayProtobuf) bool {
			if _, ok := installs[fmt.Sprint(s.GameID)]; ok {
				return true
			}
			return false
		})

		for _, g := range cache {
			if data, err := parseGameYaml([]byte(g.GameYaml)); err == nil {
				game := models.Game{
					Name:       data.Root.Name,
					Platform:   "ubisoft",
					Path:       installs[fmt.Sprint(g.GameID)],
					Executable: data.Root.StartGame.Online.Executables[0].Path.Relative,
					LaunchCmd:  "uplay://launch/" + fmt.Sprint(g.GameID) + "/1",
				}
				games = append(games, game)
			}
		}
	} else {
		// Fallback to using the registry
		data, err := findUbisoftGameData(installs)
		if err != nil {
			return nil, err
		}

		for id, name := range data {
			game := models.Game{
				Name:       name,
				Platform:   "ubisoft",
				Path:       installs[id], // Game root only
				Executable: "",           // No exe exists in the registry
				LaunchCmd:  "uplay://launch/" + id + "/1",
			}
			games = append(games, game)
		}
	}

	return games, nil
}

/*
func (s *UbisoftScanner) ScanGames() ([]db.Game, error) {
	installs, err := findUbisoftInstalls()
	if err != nil {
		return nil, err
	}

	binaries := make(map[string]string)

	// If configurations cache file exists, extract the game data.
	// We're gonna need it for getting the binaries.
	// This is an optional step so we shouldn't cancel scanner if it fails.
	uplay_games, err := parseConfigurationsFile(s.BasePath)
	if err == nil {
		uplay_games = utils.Filter(uplay_games, func(s UplayProtobuf) bool {
			if _, ok := installs[fmt.Sprint(s.GameID)]; ok {
				return true
			}
			return false
		})

		for _, g := range uplay_games {
			if data, err := parseGameYaml([]byte(g.GameYaml)); err == nil {
				binaries[fmt.Sprint(g.GameID)] = data.Root.StartGame.Online.Executables[0].Path.Relative
			}
		}

		fmt.Println(len(uplay_games))

	} else {
		fmt.Printf("ERROR parsing configurations cache file: %v", err)
	}

	var games []db.Game
	for id, data := range installs {
		game := &db.Game{
			Name:          data[0],
			Platform:      "ubisoft",
			Path:          data[1],
			Executable:    utils.TryGet(binaries, id, ""),
			LaunchOptions: "uplay://launch/" + id + "/1", // FIXME: temporary (might need a new property for online launching),
		}

		games = append(games, *game)
	}

	return games, nil
}
*/
