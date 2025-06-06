package models

type PlatformCount struct {
	Platform  string `db:"platform"`
	GameCount int    `db:"game_count"`
}
