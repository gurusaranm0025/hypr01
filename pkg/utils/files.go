package utils

import (
	"io"
	"os"
	"path/filepath"
)

func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteFile(content, path string) error {
	var err error

	if err = CreateDir(filepath.Dir(path)); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), os.ModePerm)
}

func CopyFile(sourcePath, destPath string) error {
	var info os.FileInfo
	var source, dest *os.File
	var err error

	if info, err = os.Stat(sourcePath); err != nil {
		return err
	}

	if source, err = os.Open(sourcePath); err != nil {
		return err
	}
	defer source.Close()

	if dest, err = os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode()); err != nil {
		return err
	}
	defer dest.Close()

	if _, err = io.Copy(dest, source); err != nil {
		return err
	}

	return os.Chmod(destPath, info.Mode())
}
