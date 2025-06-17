package display

import (
	"encoding/json"
	"errors"
	"gurusaranm0025/hyprone/pkg/utils"
)

type DisplaysJSON struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Width   int     `json:"width"`
	Height  int     `json:"height"`
	Scale   float64 `json:"scale"`
	Focused bool    `json:"focused"`
}

type HyprOptionJSON struct {
	Option string `json:"option"`
	Int    int    `json:"int"`
	Set    bool   `json:"set"`
}

func GetScreenresolution() (int, int, int, error) {
	var err error

	output, err := utils.ExecCommand("hyprctl -j monitors")
	if err != nil {
		return -1, -1, -1, err
	}

	var displays []DisplaysJSON
	if err = json.Unmarshal([]byte(output), &displays); err != nil {
		return -1, -1, -1, err
	}

	for _, display := range displays {
		if display.Focused {
			return display.Width, display.Height, int(display.Scale), nil
		}
	}

	return -1, -1, -1, errors.New("focused display not found")
}

func GetHyprBorder() (int, error) {
	var err error

	ouput, err := utils.ExecCommand("hyprctl -j getoption decoration:rounding")
	if err != nil {
		return -1, err
	}

	var roundingOption HyprOptionJSON
	if err = json.Unmarshal([]byte(ouput), &roundingOption); err != nil {
		return -1, err
	}

	return roundingOption.Int, nil
}
