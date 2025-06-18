package utils

import (
	"log/slog"
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

	if err = os.WriteFile(path, []byte(content), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func GetHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error(err.Error())
	}
	return home
}
