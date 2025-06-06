package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/solaire/genie/internal/db"

	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, cmd *cli.Command) error {
	if cmd.NArg() < 1 {
		return fmt.Errorf("game name or alias is required")
	}
	name := cmd.Args().First()

	game, err := db.Games().FindGameByNameOrAlias(name)
	if err != nil {
		return err
	}

	var proc *exec.Cmd

	switch {
	case game.LaunchCmd != "":
		// Launch using platform-specific command
		proc = exec.Command("cmd", "/C", "start", "", game.LaunchCmd)

	case game.Executable != "":
		// Launch directly from executable
		proc = exec.Command(game.Executable)

		// Set working dir to executable's folder
		proc.Dir = filepath.Dir(game.Executable)

	default:
		return fmt.Errorf("no valid launch method for game: %s", name)
	}

	// Optionally attach stdout/stderr for debugging
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	return proc.Run()
}
