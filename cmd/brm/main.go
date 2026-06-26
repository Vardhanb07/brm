package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:                   "brm",
		Usage:                  "Stores a file that is being deleted",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Value:   false,
				Usage:   "ignore nonexistent files and arguments, never prompt",
			},
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Value:   false,
				Usage:   "prompt before every removal",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "remove directories and their contents recursively",
			},
			&cli.BoolFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Value:   false,
				Usage:   "remove empty directories",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Value:   false,
				Usage:   "explain what is being done",
			},
			&cli.StringFlag{
				Name:    "trash",
				Aliases: []string{"t"},
				Value:   filepath.Join("~", ".local", "share", "Trash", "files"),
				Usage:   "places the delete files in trash folder",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "file",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
