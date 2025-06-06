package app

import (
	"context"
	"fmt"

	"github.com/solaire/genie/internal/db"
	"github.com/solaire/genie/pkg/logger"

	"github.com/urfave/cli/v3"
)

func PlatformList(ctx context.Context, cmd *cli.Command) error {
	platforms, err := db.Platforms().ListPlatformsWithCount()
	if err != nil {
		return err
	}
	for _, info := range platforms {
		fmt.Fprintf(logger.InfoWriter, "%s (%d)\n", info.Platform, info.GameCount)
	}
	return nil
}
