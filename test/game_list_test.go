package test

import (
	"testing"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/models"
	"github.com/solaire/genie/test/testutils"

	"github.com/stretchr/testify/assert"
)

func Test__GameList__All__NoGames(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	// Act
	err := testutils.RunCLI([]string{"game", "list"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameList__All__OneGame(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "list"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := "Witcher (platform: gog)\n"
	assert.Equal(t, want, got)
}

func Test__GameList__All__ThreeGames(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Hades", Platform: "gog", Path: `C:/Games/Hades`, Executable: `C:/Games/Hades/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Portal", Alias: "port", Platform: "steam", Path: `C:/Games/Portal`, Executable: `C:/Games/Portal/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "list"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := "Hades (platform: gog)\nPortal (alias: port) (platform: steam)\nWitcher (platform: gog)\n"
	assert.Equal(t, want, got)
}

func Test__GameList__PlatformFilter__OnePlatform_Invalid(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Hades", Platform: "gog", Path: `C:/Games/Hades`, Executable: `C:/Games/Hades/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Portal", Alias: "port", Platform: "steam", Path: `C:/Games/Portal`, Executable: `C:/Games/Portal/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "list", "--platform", "invalid"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameList__PlatformFilter__OnePlatform(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Hades", Platform: "gog", Path: `C:/Games/Hades`, Executable: `C:/Games/Hades/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Portal", Alias: "port", Platform: "steam", Path: `C:/Games/Portal`, Executable: `C:/Games/Portal/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "list", "--platform", "steam"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := "Portal (alias: port) (platform: steam)\n"
	assert.Equal(t, want, got)
}

func Test__GameList__PlatformFilter__TwoPlatforms(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Hades", Platform: "gog", Path: `C:/Games/Hades`, Executable: `C:/Games/Hades/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Portal", Alias: "port", Platform: "steam", Path: `C:/Games/Portal`, Executable: `C:/Games/Portal/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "list", "--platform", "steam, gog"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := "Hades (platform: gog)\nWitcher (platform: gog)\nPortal (alias: port) (platform: steam)\n"
	assert.Equal(t, want, got)
}
