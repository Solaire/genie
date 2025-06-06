package app

import (
	"context"
	"fmt"
	"os"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/logger"
	"github.com/solaire/genie/pkg/models"

	"github.com/urfave/cli/v3"
)

func GameList(ctx context.Context, cmd *cli.Command) error {
	var games []models.Game
	var err error
	platforms := cmd.StringSlice("platform")

	if len(platforms) == 0 {
		games, err = db.Games().ListGames()
	} else {
		games, err = db.Games().ListPlatformGames(platforms)
	}

	if err != nil {
		return err
	}

	for _, game := range games {
		if game.Name == game.Alias {
			fmt.Fprintf(logger.InfoWriter, "%s (platform: %s)\n", game.Name, game.Platform)
		} else {
			fmt.Fprintf(logger.InfoWriter, "%s (alias: %s) (platform: %s)\n", game.Name, game.Alias, game.Platform)
		}
	}
	return nil
}

func GameAdd(ctx context.Context, cmd *cli.Command) error {
	name := cmd.String("name")
	binary := cmd.String("binary")
	dir := cmd.String("dir")

	if db.Games().Exists(name) {
		return fmt.Errorf("'%s' already exists", name)
	}

	if _, err := os.Stat(dir); err != nil {
		return fmt.Errorf("invalid game dir: %s", dir)
	}

	game := &models.Game{
		Name:       name,
		Platform:   "custom",
		Path:       dir,
		Executable: binary,
	}

	return db.Games().InsertGame(game)
}

func GameRemove(ctx context.Context, cmd *cli.Command) error {
	if cmd.NArg() < 1 {
		return fmt.Errorf("game name or alias is required")
	}
	name := cmd.Args().First()

	return db.Games().DeleteGame(name)
}

func GameAlias(ctx context.Context, cmd *cli.Command) error {
	if cmd.NArg() < 1 {
		return fmt.Errorf("game name or alias is required")
	}

	name := cmd.Args().Get(0)
	alias := cmd.Args().Get(1)

	if !db.Games().Exists(name) {
		return fmt.Errorf("'%s' no such game", name)
	}

	if alias != "" {
		if existing := db.Games().AliasCheck(alias); existing != "" {
			return fmt.Errorf("'%s' alias already taken by %s", alias, existing)
		}

		return db.Games().UpdateGameAlias(name, alias)
	}

	return db.Games().DeleteGameAlias(name)
}
