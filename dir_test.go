package fs_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	fs "github.com/Vardhanb07/brm"
)

const (
	setupDirFile = "setupdir.sh"
	testDir      = "./test1"
)

func setupDir() {
	cmdPath, err := exec.LookPath("bash")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(cmdPath, setupDirFile)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func teardownDir() {
	if err := os.RemoveAll(testDir); err != nil {
		log.Fatal(err)
	}
	if err := os.RemoveAll(trashDir); err != nil {
		log.Fatal(err)
	}
}

// test dir structure
//
//	test1
//	 ├── file1.txt
//	 └── test2
//	 		 └── file2.txt
func TestRemoveDir(t *testing.T) {
	setupDir()
	defer teardownDir()
	var mockStdout bytes.Buffer
	err := fs.RemoveDir(testDir, trashDir, false, false, &mockStdout)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.ReadDir(testDir)
	if os.IsExist(err) {
		t.Errorf("expected to remove %v", testDir)
	}
	f := os.DirFS(filepath.Join(trashDir, testDir))
	if _, err := f.Open("file1.txt"); os.IsNotExist(err) {
		t.Errorf("expected to have %v in %v", filepath.Join(trashDir, testDir, "file1.txt"), trashDir)
	}
	f = os.DirFS(filepath.Join(trashDir, testDir, "test2"))
	if _, err := f.Open("file2.txt"); os.IsNotExist(err) {
		t.Errorf("expected to have %v in %v", filepath.Join(trashDir, testDir, "test2", "file2.txt"), trashDir)
	}
}

func TestRemoveDirVerbose(t *testing.T) {
	setupDir()
	defer teardownDir()
	var mockStdout bytes.Buffer
	err := fs.RemoveDir(testDir, trashDir, true, false, &mockStdout)
	if err != nil {
		t.Fatal(err)
	}
	out := make([]byte, 1024)
	n, err := mockStdout.Read(out)
	if err != nil {
		if err != io.EOF {
			t.Fatal(err)
		}
	}
	if !bytes.Contains(out[:n], []byte("descending")) {
		t.Errorf("expecting to contain desending, got: %s", out[:n])
	}
}
