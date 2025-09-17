package volume

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"strings"
)

func notifyVolume(currentVolume int) error {
	angle := int(((currentVolume + 2) / 5) * 5)

	iconPath := fmt.Sprintf("%s/.config/dunst/icons/vol/vol-%d.svg", utils.GetHomeDir(), angle)

	dotsCount := currentVolume / 15
	if dotsCount <= 0 {
		dotsCount = 1
	}
	dots := strings.Repeat(".", dotsCount)

	command := fmt.Sprintf("notify-send -a \"hyprone-controls\" -r 000001 --icon=%s %d%% %s", iconPath, currentVolume, dots)

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	return nil
}

func notifyMute(state, device string) error {
	iconPath := fmt.Sprintf("%s/.config/dunst/icons/vol/%s-%s.svg", utils.GetHomeDir(), state, device)

	var notificationID int = 0000021
	if device == "mic" {
		notificationID = 0000022
	}

	command := fmt.Sprintf("notify-send -a \"hyprone-controls\" -r %d --icon=%s %s %s", notificationID, iconPath, state, device)

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	return nil
}
