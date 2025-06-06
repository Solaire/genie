package test

import (
	"testing"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/models"
	"github.com/solaire/genie/test/testutils"

	"github.com/stretchr/testify/assert"
)

func Test__GameAlias__NoParams(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	// Act
	err := testutils.RunCLI([]string{"game", "alias"})

	// Assert
	assert.ErrorContains(t, err, "game name or alias is required")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAlias__NotFound(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Alias: "thegame", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "alias", "Hades"})

	// Assert
	assert.ErrorContains(t, err, "'Hades' no such game")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAlias__Duplicate(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Alias: "thegame", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Hades", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "alias", "Hades", "thegame"})

	// Assert
	assert.ErrorContains(t, err, "'thegame' alias already taken by Witcher")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAlias__NewAlias__CaseInsensitive(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "alias", "WITCHER", "w3"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAlias__Update__Name(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Alias: "thegame", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "alias", "Witcher", "w3"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAlias__Update__Alias(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Alias: "w3", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "alias", "w3", "thegame"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAlias__Delete__Name(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Alias: "thegame", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "alias", "Witcher"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAlias__Delete__Alias(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Alias: "w3", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "alias", "w3"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}
