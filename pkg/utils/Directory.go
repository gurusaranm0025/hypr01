package utils

import (
	"os"
	"path/filepath"
)

func WriteTo(path, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

type Entry struct {
	path string
	info os.FileInfo
}

func GetFilesAndDirs(path string, deep bool) ([]Entry, error) {
	var err error
	var entries []Entry

	if deep {
		err = filepath.Walk(path, func(p string, info os.FileInfo, f_err error) error {
			if f_err != nil {
				return f_err
			}

			if p != path {
				entries = append(entries, Entry{
					path: path,
					info: info,
				})
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		items, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, item := range items {
			info, err := item.Info()
			if err != nil {
				return nil, err
			}
			entries = append(entries, Entry{
				path: filepath.Join(path, item.Name()),
				info: info,
			})
		}
	}

	return entries, nil
}

func ListFilesAndDirs(path string, deep bool) ([]string, error) {
	var err error
	var entries []string

	if deep {
		err = filepath.Walk(path, func(p string, info os.FileInfo, f_err error) error {
			if f_err != nil {
				return f_err
			}

			if p != path {
				entries = append(entries, path)
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		items, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			entries = append(entries, filepath.Join(path, item.Name()))
		}
	}

	return entries, nil
}
