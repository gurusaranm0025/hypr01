package logout

import (
	"fmt"
	display "gurusaranm0025/hyprone/pkg/modules/Display"
	"gurusaranm0025/hyprone/pkg/utils"
	"path/filepath"
	"strings"
)

type LogoutValues map[string]string

func GetLogoutValues(layout int) (LogoutValues, error) {
	var width, height, hyprBorder, scale int
	var err error

	if width, height, scale, err = display.GetScreenresolution(); err != nil {
		return nil, err
	}

	scale = scale * 100
	if hyprBorder, err = display.GetHyprBorder(); err != nil {
		return nil, err
	}

	fontSize := fmt.Sprintf("%d", height*2/100)
	buttonRadius := fmt.Sprintf("%d", hyprBorder*8)
	activeButtonRadius := fmt.Sprintf("%d", hyprBorder*5)

	logoutValues := LogoutValues{
		"${fontSize}":             fontSize,
		"${button_radius}":        buttonRadius,
		"${active_button_radius}": activeButtonRadius,
		"${HOME}":                 utils.GetHomeDir(),
	}

	switch layout {
	case 1:
		margin := fmt.Sprintf("%d", height*28/scale)
		hover := fmt.Sprintf("%d", height*23/scale)
		logoutValues["${margin}"] = margin
		logoutValues["${hover}"] = hover
	case 2:
		x_margin := fmt.Sprintf("%d", width*35/scale)
		y_margin := fmt.Sprintf("%d", height*25/scale)
		x_hover := fmt.Sprintf("%d", width*32/scale)
		y_hover := fmt.Sprintf("%d", height*20/scale)

		logoutValues["${x_margin}"] = x_margin
		logoutValues["${y_margin}"] = y_margin
		logoutValues["${x_hover}"] = x_hover
		logoutValues["${y_hover}"] = y_hover
	}

	return logoutValues, nil
}

func Logout(layout int) error {
	var logoutValues LogoutValues
	var cols int
	var err error

	// this command is little different from other error checks
	if _, err = utils.ExecCommand("pkill wlogout"); err == nil {
		return nil
	}

	home := utils.GetHomeDir()
	layoutPath := fmt.Sprintf("%s/.config/wlogout/layout_%d", home, layout)
	stylesPath := fmt.Sprintf("%s/.config/wlogout/style_%d.css", home, layout)

	stylesContent, err := utils.ReadFile(stylesPath)
	if err != nil {
		return err
	}

	if logoutValues, err = GetLogoutValues(layout); err != nil {
		return err
	}

	for old, new := range logoutValues {
		stylesContent = strings.ReplaceAll(stylesContent, old, new)
	}

	switch layout {
	case 1:
		cols = 6
	case 2:
		cols = 2
	}

	cssPath := filepath.Join(home, ".cache/HyprOne/style.css")
	if err = utils.WriteFile(stylesContent, cssPath); err != nil {
		return err
	}

	command := fmt.Sprintf("wlogout -b %d -c 0 -r 0 -m 0 --layout %s --css %s --protocol layer-shell", cols, layoutPath, cssPath)
	if _, err = utils.ExecCommand(command); err != nil {
		return err
	}

	return nil
}
