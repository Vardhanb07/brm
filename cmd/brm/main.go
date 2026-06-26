package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	fs "github.com/Vardhanb07/brm"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:                   "brm",
		Usage:                  "Stores a file that is being deleted",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "remove directories and their contents recursively",
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
				Value:   fs.DefaultTrashDir(),
				Usage:   "places the delete files in trash folder",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "file",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			file := cmd.StringArg("file")
			recursive := cmd.Bool("recurive")
			verbose := cmd.Bool("verbose")
			trash := cmd.String("trash")
			if file == "/" || file == "/*" {
				return errors.New("brm will not delete root dir use rm instead")
			}
			fstat, err := os.Stat(file)
			if err != nil {
				return err
			}
			if fs.CheckTrashDir(trash) {
				return errors.New("trash dir does not exist")
			}
			if fstat.IsDir() && !recursive {
				return errors.New("brm will not delete a directory without -r, --recursive flag")
			}
			if !fstat.IsDir() {
				return fs.Remove(file, trash, verbose, os.Stdout)
			}
			return fs.RemoveDir(file, trash, verbose, os.Stdout)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
