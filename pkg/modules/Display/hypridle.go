package display

import (
	"gurusaranm0025/hyprone/pkg/utils"
)

func ToggleHyprIdle(val string) error {
	var err error
	var process bool

	if val == "toggle" {
		process = utils.IsProcessRunning("hypridle")
	}

	var command string
	if process || val == "0" {
		command = "pkill hypridle"
	} else if !process || val == "1" {
		command = "hyprctl dispatch exec hypridle"
	}

	_, err = utils.ExecCommand(command)

	return err
}
