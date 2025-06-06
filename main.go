package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/solaire/genie/cmd"
	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/logger"

	"github.com/urfave/cli/v3"
)

const (
	LOG_DIR string = `.local/share/genie/logs`
	APP_DIR string = `.local/share/genie/`
)

func init() {
	// Make essential directories
	app_root_dir := filepath.Join(os.Getenv("HOME"), APP_DIR)
	app_logs_dir := filepath.Join(app_root_dir, "logs")

	// Ensure app dir exists
	if err := os.MkdirAll(app_logs_dir, 0755); err != nil {
		log.Fatal(err)
	}

	// Initialise logger
	if err := logger.Init(app_logs_dir); err != nil {
		log.Fatal(err)
	}

	// Initialise database
	if err := db.Init(filepath.Join(app_root_dir, "games.db")); err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &cli.Command{
		Name:     "genie",
		Usage:    "Game scanner and launcher",
		Commands: cmd.Commands,
		Version:  "0.0.1",
	}

	logger.Printf("===== GENIE (version: %s) =====", app.Version)

	if err := app.Run(context.Background(), os.Args); err != nil {
		logger.Fatalf(err)
	}
}
