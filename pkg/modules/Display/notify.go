package display

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/conf"
	"gurusaranm0025/hyprone/pkg/utils"
	"strings"
)

func notify() error {

	percent, err := getBrightness()
	if err != nil {
		return err
	}

	angle := int(((percent + 2) / 5) * 5)

	iconPath := fmt.Sprintf("%s/.config/dunst/icons/vol/vol-%d.svg", conf.HomeDirPath, angle)

	dots := strings.Repeat(".", percent/15)

	command := fmt.Sprintf("notify-send -a \"hyprone\" -r 000004 --icon=%s %d%%%s brightness", iconPath, percent, dots)

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	return nil
}
