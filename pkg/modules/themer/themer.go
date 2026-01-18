package themer

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/common"
	"gurusaranm0025/hyprone/pkg/config"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func NewThemer(themeName string) *Themer {
	if themeName == "default" {
		themeName = common.DEFAULT_THEME
	}

	return &Themer{
		ThemeName: themeName,
	}
}

type Themer struct {
	ThemeName string
}

func (t *Themer) Install() error {
	var err error

	// CREATING DIRECTORIES
	if !config.CheckInitialSetupNE() {
		command := "hyprone --initial-setup directory"
		if _, err = utils.ExecCommand(command); err != nil {
			return err
		}
		command = "hyprone --initial-setup dependency"
		if _, err = utils.ExecCommand(command); err != nil {
			return err
		}
	}

	utils.CreateDir(common.GIT_CLONE_DIR_PATH)
	if err = os.Chdir(common.GIT_CLONE_DIR_PATH); err != nil {
		return err
	}

	// DOWNLOADING THEME
	if out, err := utils.ExecCommand(fmt.Sprintf("curl -L https://codeload.github.com/gurusaranm0025/hypr01/tar.gz/main | tar -xz --strip-components=2 hypr01-main/themes/%s", t.ThemeName)); err != nil {
		slog.Error(out)
		return err
	}

	// PLACING THE THEME IN THE CORRECT PLACE
	if err = t.filesCopier(filepath.Join(common.GIT_CLONE_DIR_PATH, t.ThemeName), utils.GetHomeDir()); err != nil {
		return err
	}

	if _, err = utils.ExecCommand("fc-cache -fv"); err != nil {
		return err
	}

	return nil
}

func (t *Themer) filesCopier(fileFolderLocation, targetLocation string) error {
	var err error
	var current_location = fileFolderLocation
	var copy_location = targetLocation

	entries, err := utils.GetFilesAndDirs(current_location, false)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		current_location = filepath.Join(fileFolderLocation, entry.Info.Name())
		copy_location = filepath.Join(targetLocation, entry.Info.Name())

		if entry.Info.IsDir() {
			if err = t.filesCopier(current_location, copy_location); err != nil {
				return err
			}
		} else {

			if err = utils.CreateDir(filepath.Dir(copy_location)); err != nil {
				return err
			}

			if strings.HasPrefix(entry.Info.Name(), "$") {
				copy_location = filepath.Join(targetLocation, strings.TrimPrefix(entry.Info.Name(), "$"))
				return t.filler(current_location, copy_location)
			} else {
				// fmt.Printf("filesCopier ===> %s --> %s\n", current_location, copy_location)
				return utils.CopyFile(current_location, copy_location)
			}
		}
	}

	return nil
}

func (t *Themer) filler(path, savePath string) error {
	// Fills the data the config asks
	file, err := utils.ReadFile(path)
	if err != nil {
		return err
	}

	for old, new := range common.PlaceholderValues {
		file = strings.ReplaceAll(file, old, new)
	}

	// fmt.Printf("filler ==> %s --> %s\n", path, savePath)
	return utils.WriteFile(file, savePath)
}
