package audio

import (
	"errors"
	"fmt"
	"gurusaranm0025/hyprone/pkg/conf"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"strconv"
	"strings"
)

func getVolume() (int, error) {
	var outInt int64
	var err error

	cmdStr := "amixer sget Master"

	output, err := utils.ExecCommand(cmdStr, true)
	if err != nil {
		return -1, err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "[on]") {
			splits := strings.Split(line, "[")
			if outInt, err = strconv.ParseInt(strings.TrimSuffix(splits[1], "%] "), 10, 0); err != nil {
				return -1, err
			}
			return int(outInt), nil
		}
	}

	return -1, errors.New("no master device, which is on, not found")
}

func changeVol(value int, mode rune) error {
	var params string
	switch mode {
	case 'i':
		params = fmt.Sprintf("%d%%+", value)
	case 'd':
		params = fmt.Sprintf("%d%%-", value)
	case 'o':
		params = fmt.Sprintf("%d%%", value)
	}

	cmdStr := fmt.Sprintf("amixer set 'Master' %s", params)
	if out, err := utils.ExecCommand(cmdStr, true); err != nil {
		slog.Info("COMMANDS OUTPUT==>")
		fmt.Println(out)
		return err
	}

	return nil
}

func notify() error {
	var dotsCount, vol int
	var err error

	if vol, err = getVolume(); err != nil {
		return err
	}

	angle := int(((vol + 2) / 5) * 5)

	iconPath := fmt.Sprintf("%s/.config/dunst/icons/vol/vol-%d.svg", conf.HomeDirPath, angle)

	if dotsCount = vol / 15; dotsCount <= 0 {
		dotsCount = 1
	}
	dots := strings.Repeat(".", dotsCount)

	notifyCMD := fmt.Sprintf("notify-send -a \"t2\" -r 911190 --icon=%s %d%% %s", iconPath, vol, dots)

	if _, err = utils.ExecCommand(notifyCMD, false); err != nil {
		return err
	}

	return nil
}

func notifyMute(mode, device string) error {
	iconPath := fmt.Sprintf("%s/.config/dunst/vol/%s-%s.svg", conf.HomeDirPath, mode, device)

	notifyCMD := fmt.Sprintf("notify-send -a \"t2\" -r 911190 --icon=%s %s", iconPath, device)

	if _, err := utils.ExecCommand(notifyCMD, false); err != nil {
		return err
	}

	return nil
}

func MuteSpeaker() error {
	var mode, out string
	var err error
	cmd := "amixer get Master"

	if out, err = utils.ExecCommand(cmd, true); err != nil {
		return err
	}

	if strings.Contains(out, "[on]") {
		cmd = "amixer set Master mute"
		mode = "muted"
	} else {
		cmd = "amixer set Master unmute"
		mode = "unmuted"
	}

	if out, err = utils.ExecCommand(cmd, true); err != nil {
		fmt.Println("OUTPUT OF THE COMMAND ==>")
		fmt.Println(out)
		return err
	}

	if err = notifyMute(mode, "speaker"); err != nil {
		return err
	}

	return nil
}

func MuteMic() error {
	var mode, out string
	var err error
	cmd := "amixer get Capture"

	if out, err = utils.ExecCommand(cmd, true); err != nil {
		return err
	}

	if strings.Contains(out, "[on]") {
		cmd = "amixer set Capture nocap"
		mode = "muted"
	} else {
		cmd = "amixer set Capture cap"
		mode = "unmuted"
	}

	if out, err = utils.ExecCommand(cmd, true); err != nil {
		fmt.Println("OUTPUT OF THE ERRORED COMMAND ==>")
		fmt.Println(out)
		return err
	}

	if err = notifyMute(mode, "mic"); err != nil {
		return err
	}

	return nil
}

func Volume(mode rune) error {
	var intVol int
	var err error

	if intVol, err = getVolume(); err != nil {
		return err
	}

	switch mode {
	case 'i':
		if intVol < 10 {
			err = changeVol(1, mode)
		} else {
			err = changeVol(3, mode)
		}
	case 'd':
		if intVol < 0 {
			err = changeVol(0, 'o')
		} else if intVol < 10 {
			err = changeVol(1, mode)
		} else {
			err = changeVol(3, mode)
		}
	}

	if err != nil {
		return err
	}

	if err = notify(); err != nil {
		return err
	}

	return nil
}
