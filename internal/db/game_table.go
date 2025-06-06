package db

import (
	"fmt"

	"github.com/solaire/genie/pkg/models"

	"github.com/jmoiron/sqlx"
)

type GameTable struct {
	db *sqlx.DB
}

func (gt *GameTable) InsertGame(game *models.Game) error {
	if game.Alias == "" {
		game.Alias = game.Name
	}

	const query = `
	INSERT INTO game
	(name, alias, platform, path, executable, launch)
	VALUES
	(:name, :alias, :platform, :path, :executable, :launch)
	`

	_, err := gt.db.NamedExec(query, game)
	return err
}

func (gt *GameTable) DeleteGame(name string) error {
	const query = `
	DELETE FROM game
	WHERE (name = ? OR alias = ?)
	AND platform = 'custom'
	`

	result, err := gt.db.Exec(query, name, name)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no game found with name or alias '%s'", name)
	}
	return nil
}

func (gt *GameTable) DeletePlatformGames(platform string) error {
	const query = `
	DELETE FROM game
	WHERE platform = ?
	`

	result, err := gt.db.Exec(query, platform)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no games found for platform: %s", platform)
	}
	return nil
}

func (gt *GameTable) DeleteAll() error {
	const query = `
	DELETE FROM game
	`

	result, err := gt.db.Exec(query)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no games to delete")
	}
	return nil
}

func (gt *GameTable) ListGames() ([]models.Game, error) {
	const query = `
	SELECT *
	FROM game
	ORDER BY name
	`

	var games []models.Game
	err := gt.db.Select(&games, query)
	return games, err
}

func (gt *GameTable) ListPlatformGames(platform_filter []string) ([]models.Game, error) {
	if len(platform_filter) == 0 {
		return nil, fmt.Errorf("must specify at least 1 platform or use ListGames()")
	}

	query := `
	SELECT *
	FROM game
	WHERE platform IN (?)
	ORDER BY platform, name
	`

	query, args, err := sqlx.In(query, platform_filter)
	if err != nil {
		return nil, err
	}

	query = gt.db.Rebind(query)
	var games []models.Game

	err = gt.db.Select(&games, query, args...)
	return games, err
}

func (gt *GameTable) FindGameByNameOrAlias(name string) (*models.Game, error) {
	const query = `
	SELECT *
	FROM game
	WHERE (name = ? OR alias = ?)
	LIMIT 1
	`

	var game models.Game
	err := gt.db.Get(&game, query, name, name)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (gt *GameTable) UpdateGameAlias(old, new string) error {
	const query = `
	UPDATE game
	SET alias = ?
	WHERE (name = ? OR alias = ?)`

	_, err := gt.db.Exec(query, new, old, old)
	return err
}

func (gt *GameTable) DeleteGameAlias(old string) error {
	return gt.UpdateGameAlias(old, old)
}

func (gt *GameTable) Exists(name string) bool {
	const query = `
	SELECT 1
	FROM game
	WHERE (name = ? OR alias = ?)`

	var val int
	gt.db.Get(&val, query, name, name)
	return val == 1
}

func (gt *GameTable) AliasCheck(alias string) string {
	const query = `
	SELECT name
	FROM game
	WHERE alias = ?`

	var name string
	gt.db.Get(&name, query, alias)
	return name
}
