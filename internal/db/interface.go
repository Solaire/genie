package db

import "github.com/solaire/genie/pkg/models"

type GameStore interface {
	InsertGame(game *models.Game) error
	DeleteGame(name string) error
	DeletePlatformGames(platform string) error
	DeleteAll() error
	ListGames() ([]models.Game, error)
	ListPlatformGames(platform_filter []string) ([]models.Game, error)
	FindGameByNameOrAlias(name string) (*models.Game, error)
	UpdateGameAlias(old, new string) error
	DeleteGameAlias(old string) error
	Exists(name string) bool
	AliasCheck(alias string) string
}

type PlatformStore interface {
	ListPlatformsWithCount() ([]models.PlatformCount, error)
}
