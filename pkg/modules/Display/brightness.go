package display

import (
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

func Brightness(mode rune) error {
	var err error

	percent, err := getBrightness()
	if err != nil {
		return err
	}

	switch mode {
	case 'i':
		if percent < 10 {
			err = setBrightness(percent + 1)
		} else {
			err = setBrightness(percent + 5)
		}
	case 'd':
		if percent < 10 {
			err = setBrightness(percent - 1)
		} else {
			err = setBrightness(percent - 5)
		}
	}

	if err != nil {
		return err
	}

	return nil
}
