package cmd

import (
	"github.com/solaire/genie/cmd/app"

	"github.com/urfave/cli/v3"
)

var Commands = []*cli.Command{
	// Scan commands
	{
		Name:  "scan",
		Usage: "Scan for installed games",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "platform",
				Aliases: []string{"p"},
				Usage:   "Limit scan to specific platform(s) e.g. --platform steam,gog",
			},
		},
		Action: app.Scan,
	},
	// Run commands
	{
		Name:   "run",
		Usage:  "Launch an installed game",
		Action: app.Run,
	},
	// Game commands
	{
		Name:  "game",
		Usage: "Manage games",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List games",
				Action:  app.GameList,
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "platform",
						Aliases: []string{"p"},
						Usage:   "Filter games to platform(s) e.g. --platform steam,gog",
					},
				},
			},
			{
				Name:   "add",
				Usage:  "Add custom game",
				Action: app.GameAdd,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Usage: "Game name", Required: true},
					&cli.StringFlag{Name: "binary", Usage: "Path to executable", Required: true},
					&cli.StringFlag{Name: "dir", Usage: "Working directory", Required: true},
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "Remove custom game",
				Action:  app.GameRemove,
			},
			{
				Name:   "alias",
				Usage:  "Add or remove alias for a game",
				Action: app.GameAlias,
			},
		},
	},
	// Platform command
	{
		Name:  "platform",
		Usage: "Manage platforms",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List installed platforms",
				Action:  app.PlatformList,
			},
		},
	},
}
