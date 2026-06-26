package fs

import (
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
func RemoveDir(dir string, trashDir string, verbose bool, w io.Writer) error {
	return nil
}
