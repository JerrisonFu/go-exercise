package stdlib

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileOperations(t *testing.T) {
	err := FileOperations()
	if err != nil {
		t.Errorf("FileOperations() error = %v", err)
	}

	if _, err := os.Stat("test.txt"); !os.IsNotExist(err) {
		t.Error("test.txt should be removed")
	}
}

func TestPathOperations(t *testing.T) {
	result := PathOperations()
	if result == "" {
		t.Error("PathOperations() should return non-empty string")
	}
}

func TestDirOperations(t *testing.T) {
	err := DirOperations()
	if err != nil {
		t.Errorf("DirOperations() error = %v", err)
	}

	if _, err := os.Stat("testdir"); !os.IsNotExist(err) {
		t.Error("testdir should be removed")
	}
}

func TestPathJoin(t *testing.T) {
	joined := filepath.Join("a", "b", "c")
	expected := filepath.Join("a", "b", "c")
	if joined != expected {
		t.Errorf("filepath.Join() = %s, want %s", joined, expected)
	}
}

func TestPathExt(t *testing.T) {
	ext := filepath.Ext("file.txt")
	if ext != ".txt" {
		t.Errorf("filepath.Ext() = %s, want .txt", ext)
	}
}
