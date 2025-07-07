package display

import (
	"gurusaranm0025/hyprone/pkg/utils"
)

func ToggleHyprIdle(val string) error {
	var err error
	process := utils.IsProcessRunning("hypridle")

	if process && val != "1" {
		_, err = utils.ExecCommand("killall -9 hypridle")
	}
	if !process && val != "0" {
		err = utils.ExecInBackground("hypridle")
	}
	return err
}
