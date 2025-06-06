package db

import (
	"fmt"
)

func ApplySchema() error {
	if database == nil {
		return fmt.Errorf("database not initialised")
	}

	if err := createGameTable(); err != nil {
		return err
	}

	return nil
}

func createGameTable() error {
	const schema = `
		CREATE TABLE IF NOT EXISTS game (
			name 			TEXT 			collate nocase,
			alias			TEXT NOT NULL 	collate nocase,
			platform 		TEXT NOT NULL 	collate nocase,
			path 			TEXT NOT NULL,
			executable 		TEXT NOT NULL,
			launch 			TEXT,

			PRIMARY KEY(name)
		);

		CREATE UNIQUE INDEX IF NOT EXISTS idx_game_alias ON game(alias);
	`

	_, err := database.Exec(schema)
	return err
}
