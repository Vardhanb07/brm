package brm_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Vardhanb07/brm"
)

const (
	testFilePath = "./whatever/test.txt"
	trashDir     = "./trash"
)

func setupFiles() {
	if err := os.MkdirAll(filepath.Dir(testFilePath), 0750); err != nil {
		log.Fatal(err)
	}
	if _, err := os.Create(testFilePath); err != nil {
		log.Fatal(err)
	}
}

func teardownFiles() {
	if err := os.RemoveAll(trashDir); err != nil {
		log.Fatal(err)
	}
	if err := os.Remove(filepath.Dir(testFilePath)); err != nil {
		log.Fatal(err)
	}
}

func TestRemove(t *testing.T) {
	setupFiles()
	defer teardownFiles()
	var mockStdout bytes.Buffer
	err := brm.Remove(testFilePath, trashDir, false, false, &mockStdout)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Open(filepath.Join(trashDir, filepath.Dir(testFilePath), filepath.Base(testFilePath)))
	if err != nil {
		if os.IsNotExist(err) {
			t.Errorf("expected: %v, got: %v", filepath.Join(trashDir, filepath.Dir(testFilePath), filepath.Base(testFilePath)), trashDir)
		} else {
			t.Fatal(err)
		}
	}
	_, err = os.Open(testFilePath)
	if !os.IsNotExist(err) {
		t.Errorf("expected to delete file %v", testFilePath)
	}
}

func TestRemoveVerbose(t *testing.T) {
	setupFiles()
	defer teardownFiles()
	var mockStdout bytes.Buffer
	err := brm.Remove(testFilePath, trashDir, true, false, &mockStdout)
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
	if !bytes.Contains(out[:n], []byte("copying")) || !bytes.Contains(out[:n], []byte("removing")) {
		t.Errorf("expected to have copying and remove, got: %s", out[:n])
	}
}

func TestRemoveVerboseNoSave(t *testing.T) {
	setupFiles()
	defer teardownFiles()
	var mockStdout bytes.Buffer
	err := brm.Remove(testFilePath, trashDir, true, true, &mockStdout)
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
	if bytes.Contains(out[:n], []byte("copying")) {
		t.Errorf("expected to not have copying, got: %s", out[:n])
	}
}

func TestRemoveNoSave(t *testing.T) {
	setupFiles()
	defer teardownFiles()
	var mockStdout bytes.Buffer
	err := brm.Remove(testFilePath, trashDir, false, true, &mockStdout)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.ReadDir(filepath.Join(trashDir, filepath.Dir(testFilePath))); os.IsExist(err) {
		t.Errorf("expected to remove %s", filepath.Join(trashDir, filepath.Dir(testFilePath)))
	}
	if _, err := os.ReadFile(filepath.Join(trashDir, testFilePath)); os.IsExist(err) {
		t.Errorf("expected to remove %s", filepath.Join(trashDir, testFilePath))
	}
}
