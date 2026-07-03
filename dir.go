package brm

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// gets default trash dir
func DefaultTrashDir() string {
	hdir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(hdir, ".local", "share", "Trash", "files")
}

// checks if trash dir exits
func CheckTrashDir(trashDir string) bool {
	_, err := os.ReadDir(trashDir)
	return os.IsExist(err)
}

// removes dir and it's contents
func RemoveDir(dir string, trashDir string, verbose bool, noSave bool, out io.Writer) error {
	fstat, err := os.Stat(dir)
	if err != nil {
		return err
	}
	if !fstat.IsDir() {
		return Remove(dir, trashDir, verbose, noSave, out)
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		nextDir := filepath.Join(dir, file.Name())
		if !file.IsDir() {
			err := Remove(nextDir, trashDir, verbose, noSave, out)
			if err != nil {
				return err
			}
		} else {
			if verbose {
				fmt.Fprintf(out, "descending to %v driectory\n", nextDir)
			}
			err := RemoveDir(nextDir, trashDir, verbose, noSave, out)
			if err != nil {
				return err
			}
		}
	}
	if verbose {
		fmt.Fprintf(out, "removing %v directory\n", dir)
	}
	return os.Remove(dir)
}
