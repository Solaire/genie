package db

import (
	"github.com/solaire/genie/pkg/models"

	"github.com/jmoiron/sqlx"
)

type PlatformTable struct {
	db *sqlx.DB
}

func (pt *PlatformTable) ListPlatformsWithCount() ([]models.PlatformCount, error) {
	const query = `
	SELECT 
		platform,
		count(*) AS game_count
	FROM game
	GROUP BY platform
	ORDER BY platform
	`

	var platforms []models.PlatformCount
	err := pt.db.Select(&platforms, query)
	return platforms, err
}
