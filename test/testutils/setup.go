package testutils

import (
	"bytes"
	"context"
	"testing"

	"github.com/solaire/genie/cmd"
	"github.com/solaire/genie/pkg/logger"

	"github.com/urfave/cli/v3"
)

func Setup(t *testing.T) (*bytes.Buffer, func()) {
	// Set up in-memory test database.
	test_db := NewTestDB(t)

	// Capture the logger output.
	var buf bytes.Buffer
	logger.InfoWriter = &buf

	return &buf, func() {
		test_db.Close()
		logger.InfoWriter = nil
	}
}

func RunCLI(args []string) error {
	app := &cli.Command{
		Name:     "genie_test",
		Usage:    "Game scanner and launcher",
		Commands: cmd.Commands,
	}

	args = append([]string{"genie_test"}, args...)

	return app.Run(context.Background(), args)
}
