package models

type Game struct {
	Name       string `db:"name"`
	Alias      string `db:"alias"`
	Platform   string `db:"platform"`
	Path       string `db:"path"`
	Executable string `db:"executable"`
	LaunchCmd  string `db:"launch"`
}
