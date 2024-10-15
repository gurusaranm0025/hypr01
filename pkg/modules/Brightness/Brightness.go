package brightness

import (
	"errors"
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var GETBRIGHTNESS = "brightnessctl -m"
var GETBACKLIGHTNAME = "brightnessctl info"

func regexBuilderAndMatcher(pattern, source string) string {
	re := regexp.MustCompile(pattern)
	return re.FindString(source)
}

func getBrightness() (int, string, error) {
	var outInt int64
	var outStr string
	var err error

	if outStr, err = utils.ExecCommand(GETBRIGHTNESS, true); err != nil {
		return -1, "", err
	}

	outStr = regexBuilderAndMatcher("[0-9]+%", outStr)
	if outInt, err = strconv.ParseInt(outStr[:len(outStr)-1], 10, 0); err != nil {
		return -1, "", err
	}

	return int(outInt), outStr, nil
}

func getBacklightName() (string, error) {
	out, err := utils.ExecCommand(GETBACKLIGHTNAME, true)
	if err != nil {
		return "", err
	}

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Device") {
			splits := strings.Split(line, "'")
			if len(splits) >= 3 {
				return splits[1], nil
			}
		}
	}
	return "", errors.New("device not found")
}

func notify() error {
	var intBright, dotsCount int
	var backlightName, homeDir string
	var err error

	if intBright, _, err = getBrightness(); err != nil {
		return err
	}

	if backlightName, err = getBacklightName(); err != nil {
		return err
	}

	angle := int(((intBright + 2) / 5) * 5)
	if homeDir, err = os.UserHomeDir(); err != nil {
		return err
	}

	iconPath := fmt.Sprintf("%s/.config/dunst/icons/vol/vol-%d.svg", homeDir, angle)
	if dotsCount = intBright / 15; dotsCount < 0 {
		dotsCount = 0
	}
	dots := strings.Repeat(".", dotsCount)

	notifyCmd := fmt.Sprintf("notify-send -a \"t2\" -r 91190 --icon=%s %d%%%s %s", iconPath, intBright, dots, backlightName)
	if _, err = utils.ExecCommand(notifyCmd, false); err != nil {
		return err
	}
	return nil
}

func changeBrightness(value int, mode rune) error {
	var modeSym string
	if mode == 'i' {
		modeSym = "+"
	} else if mode == 'd' {
		modeSym = "-"
	} else {
		modeSym = ""
	}

	cmd := fmt.Sprintf("brightnessctl set %d%%%s", value, modeSym)
	if _, err := utils.ExecCommand(cmd, false); err != nil {
		return err
	}
	if err := notify(); err != nil {
		return err
	}
	return nil
}

func Brightness(mode rune) error {
	var intBright int
	var err error

	if intBright, _, err = getBrightness(); err != nil {
		return nil
	}

	switch mode {
	case 'i':
		if intBright < 10 {
			if err = changeBrightness(1, mode); err != nil {
				return err
			}
		} else {
			if err = changeBrightness(5, mode); err != nil {
				return err
			}
		}
	case 'd':
		if intBright < 1 {

			if err = changeBrightness(1, ' '); err != nil {
				return err
			}
		} else if intBright < 10 {
			if err = changeBrightness(1, 'd'); err != nil {
				return err
			}
		} else {
			if err = changeBrightness(5, 'd'); err != nil {
				return err
			}
		}
	}
	return nil
}
