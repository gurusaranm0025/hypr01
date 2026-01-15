package utils

import (
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

	dirPath := filepath.Dir(path)
	info, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err = CreateDir(dirPath); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if !info.IsDir() {
		if err = CreateDir(dirPath); err != nil {
			return err
		}
	}

	if err = os.WriteFile(path, []byte(content), os.ModePerm); err != nil {
		return err
	}

	return nil
}
