package wallapaper

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/common"
	"gurusaranm0025/hyprone/pkg/utils"
)

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
