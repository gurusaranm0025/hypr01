package display

import (
	"gurusaranm0025/hyprone/pkg/utils"
)

func ToggleHyprIdle(val string) error {
	var err error

	if val == "toggle" {
		if process := utils.IsProcessRunning("hypridle"); process {
			_, err = utils.ExecCommand("killall -9 hypridle")
		} else {
			err = utils.ExecInBackground("hypridle")
		}
		return err
	}

	if val == "0" {
		_, err = utils.ExecCommand("killall -9 hypridle")
	} else {
		err = utils.ExecInBackground("hypridle")
	}

	return err
}
