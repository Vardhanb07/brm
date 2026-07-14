package brm

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type file struct {
	name string
	dir  string
}

// removes file and places it in trash directory
func Remove(fpath string, trashdir string, verbose bool, noSave bool, out io.Writer) error {
	file := file{
		name: filepath.Base(fpath),
		dir:  filepath.Join(trashdir, filepath.Dir(fpath)),
	}
	if !noSave {
		if err := os.MkdirAll(file.dir, 0750); err != nil {
			return err
		}
		data, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(file.dir, file.name))
		if err != nil {
			return err
		}
		f.Write(data)
		defer f.Close()
	}
	if verbose {
		if !noSave {
			fmt.Fprintf(out, "copying %v to %v\n", fpath, filepath.Join(file.dir, file.name))
		}
		fmt.Fprintf(out, "removing %v\n", fpath)
	}
	return os.Remove(fpath)
}

func GetFileOrLinkStats(path string) (os.FileInfo, error) {
	stat, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			stat, err = os.Stat(path)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}
	return stat, nil
}
