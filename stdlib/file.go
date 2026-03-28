package stdlib

import (
	"io"
	"os"
	"path/filepath"
)

func FileOperations() error {
	filename := "test.txt"

	err := os.WriteFile(filename, []byte("Hello, Go!"), 0644)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	_ = string(data)

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}

func PathOperations() string {
	dir := "/home/user/documents"
	file := "report.pdf"

	joined := filepath.Join(dir, file)
	ext := filepath.Ext(file)
	base := filepath.Base(file)
	dirname := filepath.Dir(file)

	return joined + "|" + ext + "|" + base + "|" + dirname
}

func DirOperations() error {
	dir := "testdir"
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	_ = len(entries)

	return os.Remove(dir)
}

func CopyData() error {
	src := "source.txt"
	dst := "dest.txt"

	err := os.WriteFile(src, []byte("data"), 0644)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	os.Remove(src)
	os.Remove(dst)

	return nil
}
