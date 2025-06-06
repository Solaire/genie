package testutils

import (
	"testing"

	"github.com/solaire/genie/internal/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func NewTestDB(t *testing.T) *sqlx.DB {
	t.Helper()

	mem_db, err := sqlx.Connect("sqlite3", ":memory:")
	require.NoError(t, err)

	db.SetDB(mem_db)

	err = db.ApplySchema()
	require.NoError(t, err)

	t.Cleanup(func() {
		mem_db.Close()
	})

	return mem_db
}
