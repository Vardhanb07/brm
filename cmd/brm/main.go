package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/Vardhanb07/brm"
	"github.com/urfave/cli/v3"
)

func checkArch() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

func main() {
	cmd := &cli.Command{
		Name:                   "brm",
		Usage:                  "Stores a file that is being deleted",
		Version:                "v1.0.3",
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
				Aliases: []string{"V"},
				Value:   false,
				Usage:   "explain what is being done",
			},
			&cli.StringFlag{
				Name:    "trash",
				Aliases: []string{"t"},
				Value:   brm.DefaultTrashDir(),
				Usage:   "places the delete files in trash folder",
			},
			&cli.BoolFlag{
				Name:    "no-save",
				Aliases: []string{"n"},
				Value:   false,
				Usage:   "removes files without saving them",
			},
			&cli.BoolFlag{
				Name:    "update",
				Aliases: []string{"u"},
				Value:   false,
				Usage:   "updates brm to its latest version",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArgs{
				Name: "files",
				Min:  0,
				Max:  -1,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if !checkArch() {
				return errors.New("platform not supported")
			}
			files := cmd.StringArgs("files")
			recursive := cmd.Bool("recursive")
			verbose := cmd.Bool("verbose")
			trash := cmd.String("trash")
			noSave := cmd.Bool("no-save")
			update := cmd.Bool("update")
			switch {
			case update:
				return brm.Update(verbose, os.Stdout)
			default:
				for _, file := range files {
					resolved, err := brm.PathResolve(file)
					if err != nil {
						return err
					}
					fstat, err := brm.GetFileOrLinkStats(resolved)
					if err != nil {
						return err
					}
					if brm.CheckTrashDir(trash) {
						return errors.New("trash dir does not exist")
					}
					if fstat.IsDir() && !recursive {
						return errors.New("brm will not delete a directory without -r, --recursive flag")
					}
					if !fstat.IsDir() {
						if err := brm.Remove(file, trash, verbose, noSave, os.Stdout); err != nil {
							return err
						}
					} else if err := brm.RemoveDir(file, trash, verbose, noSave, os.Stdout); err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "brm: %v\n", err)
		os.Exit(1)
	}
}
