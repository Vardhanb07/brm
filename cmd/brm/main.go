package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"
	"slices"

	fs "github.com/Vardhanb07/brm"
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
		Version:                "v1.0.0",
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
				Aliases: []string{"vb", "e"},
				Value:   false,
				Usage:   "explain what is being done",
			},
			&cli.StringFlag{
				Name:    "trash",
				Aliases: []string{"t"},
				Value:   fs.DefaultTrashDir(),
				Usage:   "places the delete files in trash folder",
			},
			&cli.BoolFlag{
				Name:    "no-save",
				Aliases: []string{"n"},
				Value:   false,
				Usage:   "removes files without saving them",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArgs{
				Name: "files",
				Min:  1,
				Max:  -1,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cli.VersionPrinter = func(cmd *cli.Command) {
				fmt.Fprintf(os.Stdout, "version=%s\n", cmd.Root().Version)
			}
			if !checkArch() {
				return errors.New("platform not supported")
			}
			files := cmd.StringArgs("files")
			recursive := cmd.Bool("recursive")
			verbose := cmd.Bool("verbose")
			trash := cmd.String("trash")
			noSave := cmd.Bool("no-save")
			if slices.Contains(files, "/") || slices.Contains(files, "/*") {
				return errors.New("brm will not delete root dir use rm instead")
			}
			for _, file := range files {
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
					if err := fs.Remove(file, trash, verbose, noSave, os.Stdout); err != nil {
						return err
					}
				} else if err := fs.RemoveDir(file, trash, verbose, noSave, os.Stdout); err != nil {
					return err
				}
			}
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
