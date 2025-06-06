package test

import (
	"testing"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/models"
	"github.com/solaire/genie/test/testutils"

	"github.com/stretchr/testify/assert"
)

func Test__GameAdd__NoFlags(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "add"})

	// Assert
	assert.ErrorContains(t, err, `Required flags "name, binary, dir" not set`)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAdd__NoName(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	// Act
	err := testutils.RunCLI([]string{"game", "add", "--binary", "C:/Games/Witcher/game.exe", "--dir", "C:/Games/Witcher"})

	// Assert
	assert.ErrorContains(t, err, `Required flag "name" not set`)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAdd__NoBinary(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	// Act
	err := testutils.RunCLI([]string{"game", "add", "--name", "Witcher", "--dir", "C:/Games/Witcher"})

	// Assert
	assert.ErrorContains(t, err, `Required flag "binary" not set`)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAdd__NoDir(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "add", "--name", "WITCHER", "--binary", "C:/Games/Witcher/game.exe"})

	// Assert
	assert.ErrorContains(t, err, `Required flags "dir" not set`)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAdd__Duplicate__CaseInsensitive(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "add", "--name", "WITCHER", "--binary", "C:/Games/Witcher/game.exe", "--dir", "C:/Games/Witcher"})

	// Assert
	assert.ErrorContains(t, err, "'WITCHER' already exists")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAdd__DirNotExist(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "add", "--name", "New", "--binary", "C:/Games/Witcher/game.exe", "--dir", "A:/invalid"})

	// Assert
	assert.ErrorContains(t, err, "invalid game dir: A:/invalid")

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}

func Test__GameAdd__Success(t *testing.T) {
	// Arrange
	buffer, fn := testutils.Setup(t)
	t.Cleanup(fn)

	db.Games().InsertGame(&models.Game{Name: "Witcher", Platform: "gog", Path: `C:/Games/Witcher`, Executable: `C:/Games/Witcher/game.exe`})

	// Act
	err := testutils.RunCLI([]string{"game", "add", "--name", "Hades", "--binary", "C:/Games/Hades/game.exe", "--dir", "C:"})

	// Assert
	assert.NoError(t, err)

	got := buffer.String()
	want := ""
	assert.Equal(t, want, got)
}
