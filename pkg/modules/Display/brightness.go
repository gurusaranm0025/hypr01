package display

import (
	"errors"
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

func getBrightness() (int, error) {
	var output string
	var err error

	if output, err = utils.ExecCommand("brightnessctl m"); err != nil {
		return -1, err
	}
	maxBrightness, err := strconv.Atoi(strings.TrimSpace(output))
	if err != nil {
		return -1, err
	}

	if output, err = utils.ExecCommand("brightnessctl g"); err != nil {
		return -1, err
	}
	currentBrightness, err := strconv.Atoi(strings.TrimSpace(output))
	if err != nil {
		return -1, err
	}

	percent := float64(currentBrightness) / float64(maxBrightness) * 100
	return int(math.Round(percent)), nil
}

func setBrightness(val int) error {
	val = max(min(val, 100), 0)

	command := fmt.Sprintf("brightnessctl s %d%%", val)

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	if err := notify(); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func Brightness(value string) error {
	var percent int
	var err error

	if percent, err = getBrightness(); err != nil {
		return err
	}

	switch {
	case value == "+":
		if percent < 10 {
			err = setBrightness(percent + 1)
		} else {
			err = setBrightness(percent + 5)
		}
	case value == "-":
		if percent < 10 {
			err = setBrightness(percent - 1)
		} else {
			err = setBrightness(percent - 5)
		}
	case strings.HasSuffix(value, "%+") && len(value) > 2:
		var changeValue int
		value = strings.TrimSuffix(value, "%+")
		changeValue, err = strconv.Atoi(value)
		if err != nil {
			return err
		}
		err = setBrightness(percent + changeValue)
	case strings.HasSuffix(value, "%-") && len(value) > 2:
		var changeValue int
		value = strings.TrimSuffix(value, "%-")
		changeValue, err = strconv.Atoi(value)
		if err != nil {
			return err
		}
		err = setBrightness(percent - changeValue)
	case strings.HasSuffix(value, "%") && len(value) > 1:
		value = strings.TrimSuffix(value, "%")
		if percent, err = strconv.Atoi(value); err != nil {
			return err
		}
		err = setBrightness(percent)
	default:
		err = errors.New("invalid input")
	}

	if err != nil {
		return err
	}

	return nil
}
