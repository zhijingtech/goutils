package main

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOldFiles(t *testing.T) {
	dir := t.TempDir()
	oldFile := filepath.Join(dir, "old.txt")
	youngFile := filepath.Join(dir, "young.txt")
	err := os.WriteFile(oldFile, []byte("old"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	err = os.WriteFile(youngFile, []byte("young"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = RemoveOldFiles(dir, time.Millisecond*500)
	if err != nil {
		t.Fatal(err)
	}
	var fileNames []string
	err = filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(fileNames) != 1 || fileNames[0] != youngFile {
		t.Errorf("expected only young file, got %v", fileNames)
	}
}

func TestCreateAndBackupFile(t *testing.T) {
	dir := t.TempDir()
	file := path.Join(dir, "TestCreateAndBackupFile")
	f, err := CreateAndBackupFile(file)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("CreateAndBackupFile() file not exist")
	}
	_, _ = f.Write([]byte("123"))
	f.Close()
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("CreateAndBackupFile() file not exist")
	}
	f, err = CreateAndBackupFile(file)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("CreateAndBackupFile() file not exist")
	}
	_, _ = f.Write([]byte("123"))
	f.Close()
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("CreateAndBackupFile() file not exist")
	}

	var fileNames []string
	err = filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, fileNames, 2)
	assert.Contains(t, fileNames, file)
}
