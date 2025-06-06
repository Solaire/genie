package test

import (
	"testing"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/models"
	"github.com/solaire/genie/test/testutils"

	"github.com/stretchr/testify/assert"
)

func Test__GameRemove__NoParams(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	// Act
	err := testutils.RunCLI([]string{"game", "remove"})

	// Assert
	assert.ErrorContains(t, err, "game name or alias is required")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameRemove__NotFound(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "remove", "Hades"})

	// Assert
	assert.ErrorContains(t, err, "no game found with name or alias 'Hades'")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameRemove__NotCustom(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "remove", "Witcher"})

	// Assert
	assert.ErrorContains(t, err, "no game found with name or alias 'Witcher'")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameRemove__Success__Name(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "remove", "Witcher"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameRemove__Success__Alias(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Alias: "w3", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "remove", "w3"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameRemove__Success__CaseInsensitive(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "custom", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "remove", "WITCHER"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}
