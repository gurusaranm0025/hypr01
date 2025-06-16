package volume

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"strconv"
	"strings"
)

var device_ids map[string]string = map[string]string{
	"mic":     "@DEFAULT_SOURCE@",
	"speaker": "@DEFAULT_SINK@",
}

func getVolume(device_id string) (int, error) {
	command := fmt.Sprintf("wpctl get-volume %s", device_ids[device_id])

	output, err := utils.ExecCommand(command)
	if err != nil {
		return -1, err
	}
	output = strings.Split(output, ":")[1]
	if strings.Contains(output, "[MUTED]") {
		return -999, nil
	}
	output = strings.TrimSpace(output)
	volumeVal, err := strconv.ParseFloat(output, 64)
	if err != nil {
		return -1, err
	}

	return int(volumeVal * 100), nil
}

func setVolume(device_id string, val int) error {
	// clamps the valume value between 0 and 100
	val = min(max(val, 0), 100)

	command := fmt.Sprintf("wpctl set-volume %s %d%%", device_ids[device_id], val)

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	if err := notifyVolume(val); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func Mute(device_id string) error {
	command := fmt.Sprintf("wpctl set-mute %s toggle", device_ids[device_id])

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	currentVolume, err := getVolume(device_id)
	if err != nil {
		return err
	}

	var state string
	if currentVolume == -999 {
		state = "muted"
	} else {
		state = "unmuted"
	}

	if err := notifyMute(state, device_id); err != nil {
		return err
	}

	return nil

}

func Volume(mode rune) error {
	var currentVolume int
	var err error

	if currentVolume, err = getVolume("speaker"); err != nil {
		return err
	}

	switch mode {
	case 'i':
		if currentVolume < 10 {
			err = setVolume("speaker", currentVolume+1)
		} else {
			err = setVolume("speaker", currentVolume+3)
		}
	case 'd':
		if currentVolume <= 0 {
			err = setVolume("speaker", 0)
		} else if currentVolume < 10 {
			err = setVolume("speaker", currentVolume-1)
		} else {
			err = setVolume("speaker", currentVolume-3)
		}
	}

	if err != nil {
		return err
	}

	return nil
}
