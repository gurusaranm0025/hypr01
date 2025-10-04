package volume

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
)

var vol_notify_icons map[string]string = map[string]string{
	"low":    "audio-volume-low",
	"medium": "audio-volume-medium",
	"high":   "audio-volume-high",
	"muted":  "audio-volume-muted",
}

func getVolumeIcon(currentVolume int) string {
	if currentVolume == 0 {
		return vol_notify_icons["muted"]
	}
	if currentVolume <= 30 {
		return vol_notify_icons["low"]
	}
	if currentVolume <= 70 {
		return vol_notify_icons["medium"]
	}
	return vol_notify_icons["high"]
}

func notifyVolume(currentVolume int) error {
	// angle := int(((currentVolume + 2) / 5) * 5)

	// iconPath := fmt.Sprintf("%s/.config/dunst/icons/vol/vol-%d.svg", utils.GetHomeDir(), angle)

	// dotsCount := currentVolume / 15
	// if dotsCount <= 0 {
	// 	dotsCount = 1
	// }
	// dots := strings.Repeat(".", dotsCount)

	// command := fmt.Sprintf("notify-send -a \"hyprone-controls\" --transient -r 000001 --icon=%s %d%% %s", iconPath, currentVolume, dots)

	command := fmt.Sprintf("swayosd-client --custom-icon=\"%s\" --custom-progress=%.2f", getVolumeIcon(currentVolume), float32(currentVolume)/float32(100))

	if ouput, err := utils.ExecCommand(command); err != nil {
		slog.Error(ouput)
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

	command := fmt.Sprintf("notify-send -a \"hyprone-controls\" --transient -r %d --icon=%s %s %s", notificationID, iconPath, state, device)

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	return nil
}
