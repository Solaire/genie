package test

import (
	"testing"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/models"
	"github.com/solaire/genie/test/testutils"

	"github.com/stretchr/testify/assert"
)

func Test__PlatformList__NoPlatforms(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	// Act
	err := testutils.RunCLI([]string{"platform", "list"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__PlatformList__OnePlatform(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"platform", "list"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := "gog (1)\n"
	assert.Equal(t, want, got)
}

func Test__PlatformList__TwoPlatforms(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})
	db.Games().InsertGame(&models.Game{Name: "Hades", Platform: "gog", Path: `C:/Games/Hades`, Executable: `C:/Games/Witcher/Hades.exe`})
	db.Games().InsertGame(&models.Game{Name: "Portal", Platform: "steam", Path: `C:/Games/Portal`, Executable: `C:/Games/Witcher/Portal.exe`})

	// Act
	err := testutils.RunCLI([]string{"platform", "list"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := "gog (2)\nsteam (1)\n"
	assert.Equal(t, want, got)
}
