package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var database *sqlx.DB
var gameStore GameStore
var platformStore PlatformStore

func Init(db_path string) error {
	db, err := sqlx.Connect("sqlite3", db_path)
	if err != nil {
		return err
	}

	SetDB(db)

	if err := ApplySchema(); err != nil {
		return err
	}

	return nil
}

func SetDB(db *sqlx.DB) {
	database = db
	gameStore = &GameTable{database}
	platformStore = &PlatformTable{database}
}

func Games() GameStore {
	return gameStore
}

func Platforms() PlatformStore {
	return platformStore
}
