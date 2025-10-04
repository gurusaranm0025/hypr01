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
	if currentVolume == -999 {
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
	command := fmt.Sprintf("swayosd-client --custom-icon=\"%s\" --custom-progress=%.2f", getVolumeIcon(currentVolume), float32(currentVolume)/float32(100))

	if ouput, err := utils.ExecCommand(command); err != nil {
		slog.Error(ouput)
		return err
	}

	return nil
}

func notifyMute(currentVolume int, device string) error {
	// fmt.Println("==> ", currentVolume, " ==> device : ", device)
	// fmt.Println("==> ", getVolumeIcon(currentVolume))
	// fmt.Println("==> ", float32(currentVolume)/float32(100))
	var command string
	if device == "mic" {
		command = fmt.Sprintf("swayosd-client --custom-icon=\"%s\" --custom-progress=%.2f --custom-progress-text=\"mic\"", getVolumeIcon(currentVolume), float32(currentVolume)/float32(100))
	} else {
		command = fmt.Sprintf("swayosd-client --custom-icon=\"%s\" --custom-progress=%.2f", getVolumeIcon(currentVolume), max(float32(currentVolume)/float32(100), 0.01))
	}

	if output, err := utils.ExecCommand(command); err != nil {
		slog.Error(output)
		return err
	}

	return nil
}
