package wallapaper

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/common"
	"gurusaranm0025/hyprone/pkg/utils"
)

// var WALLPAPERS_FOLDER = filepath.Join(utils.GetHomeDir(), "/.HyprOne/Walls")
// var WALLPAPER_SELECTOR_ROFI_CONFIG_PATH = filepath.Join(utils.GetHomeDir(), "/.config/rofi/wallpapers/config.rasi")

func WallpaperGUI() error {
	command := fmt.Sprintf("%s/wallpaper_selector.sh", common.SCRIPTS_DIR_PATH)

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

	if err = utils.CreateDir(common.ALL_WALLS_DIR_PATH); err != nil {
		return err
	}

	return nil
}

func SetHyprlockWall() error {
	return nil
}
