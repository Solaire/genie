package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/internal/utils"
	"github.com/solaire/genie/pkg/logger"
	"github.com/solaire/genie/pkg/models"
	"github.com/solaire/genie/pkg/scanner"

	"github.com/urfave/cli/v3"
)

type _ScanResult struct {
	Platform string
	Games    []models.Game
}

func Scan(ctx context.Context, cmd *cli.Command) error {
	statuses := &scanner.ScanStatus{
		Lines: make(map[string]int),
	}

	platforms := filterPlatforms(cmd.StringSlice("platform"))
	for i, p := range platforms {
		statuses.LineInit(i, p.Name(), p.Name()+": Starting...")
	}

	var wg sync.WaitGroup
	results := make(chan _ScanResult)

	for _, platform := range platforms {

		wg.Add(1)
		go func(platform_scanner scanner.Scanner) {
			defer wg.Done()
			name := platform_scanner.Name()

			statuses.Set(name, fmt.Sprintf("%s: Detecting...", name))
			if !platform_scanner.Detect() {
				statuses.Set(name, fmt.Sprintf("\033[33m%s: Not found\033[0m", name)) // Yellow
				return
			}

			statuses.Set(name, fmt.Sprintf("%s: Scanning games...", name))
			games, err := platform_scanner.ScanGames()
			if err != nil {
				statuses.Set(name, fmt.Sprintf("\033[31m%s: Error - %v\033[0m", name, err)) // Red
				return
			}

			statuses.Set(name, fmt.Sprintf("\033[32m%s: Found %d games\033[0m", name, len(games))) // Green
			results <- _ScanResult{platform_scanner.Name(), games}
		}(platform)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	found_count := 0
	new_count := 0
	delete_count := 0

	for res := range results {
		current, err := getCurrentGames(res.Platform)
		if err != nil {
			logger.Errorf("Could not retrieve games for %s: %v\n", res.Platform, err)
			continue
		}

		found_count += len(res.Games)

		for _, game := range res.Games {
			if utils.ExistsInMap(current, game.Name) {
				logger.Printf("Game %s already exists. Skipping...\n", game.Name)
				delete(current, game.Name)
				continue
			}

			db.Games().InsertGame(&game)
			logger.Printf("New %s game: %s\n", game.Platform, game.Name)
			new_count++
		}

		// Delete unmatched games
		for name, obj := range current {
			db.Games().DeleteGame(name)
			logger.Printf("Removed %s game: %s\n", obj.Platform, name)
			delete_count++
		}
	}

	fmt.Println()
	logger.InfoWritef("Found %d games\n", found_count)
	logger.InfoWritef("Added %d games\n", new_count)
	logger.InfoWritef("Removed %d games\n", delete_count)

	return nil
}

func filterPlatforms(cli_platforms []string) []scanner.Scanner {
	if len(cli_platforms) == 0 {
		return scanner.ScannerList
	}

	var scanners []scanner.Scanner

	for _, s := range scanner.ScannerList {
		if utils.ExistsInArray(cli_platforms, s.Name()) {
			scanners = append(scanners, s)
		}
	}

	return scanners
}

func getCurrentGames(platform string) (map[string]models.Game, error) {
	game_map := make(map[string]models.Game)

	current, err := db.Games().ListPlatformGames([]string{platform})
	if err != nil {
		return nil, err
	}

	for _, game := range current {
		game_map[game.Name] = game
	}

	return game_map, nil
}
