package wallapaper

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"path/filepath"
)

var WALLPAPERS_FOLDER = filepath.Join(utils.GetHomeDir(), "/.HyprOne/Walls")

func WallpaperGUI() error {
	command := fmt.Sprintf("waypaper --folder %s", WALLPAPERS_FOLDER)

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}
	return nil
}

func StartDaemon() error {
	var err error

	utils.ExecCommand("killall -9 swww-daemon")
	if err = utils.ExecInBackground("swww-daemon"); err != nil {
		return err
	}

	if err = utils.CreateDir(WALLPAPERS_FOLDER); err != nil {
		return err
	}

	return nil
}
