package logout

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/conf"
	display "gurusaranm0025/hyprone/pkg/modules/Display"
	"gurusaranm0025/hyprone/pkg/utils"
	"path/filepath"
	"strings"
)

func Logout(layout int) error {
	var width, height, hyprBorder, cols, scale int
	var err error

	if width, height, scale, err = display.GetScreenresolution(); err != nil {
		return err
	}

	scale = scale * 100
	if hyprBorder, err = display.GetHyprBorder(); err != nil {
		return err
	}

	layoutPath := fmt.Sprintf("%s/.config/wlogout/layout_%d", conf.HomeDirPath, layout)
	stylesPath := fmt.Sprintf("%s/.config/wlogout/style_%d.css", conf.HomeDirPath, layout)

	stylesContent, err := utils.ReadFile(stylesPath)
	if err != nil {
		return err
	}

	fontSize := fmt.Sprintf("%d", height*2/100)
	buttonRadius := fmt.Sprintf("%d", hyprBorder*8)
	activeButtonRadius := fmt.Sprintf("%d", hyprBorder*5)

	stylesContent = strings.ReplaceAll(stylesContent, "${fontSize}", fontSize)
	stylesContent = strings.ReplaceAll(stylesContent, "${active_button_radius}", activeButtonRadius)
	stylesContent = strings.ReplaceAll(stylesContent, "${button_radius}", buttonRadius)
	stylesContent = strings.ReplaceAll(stylesContent, "${HOME}", conf.HomeDirPath)

	switch layout {
	case 1:
		cols = 6
		margin := fmt.Sprintf("%d", height*28/scale)
		hover := fmt.Sprintf("%d", height*23/scale)
		stylesContent = strings.ReplaceAll(stylesContent, "${margin}", margin)
		stylesContent = strings.ReplaceAll(stylesContent, "${hover}", hover)
	case 2:
		cols = 2
		x_margin := fmt.Sprintf("%d", width*35/scale)
		y_margin := fmt.Sprintf("%d", height*25/scale)
		x_hover := fmt.Sprintf("%d", width*32/scale)
		y_hover := fmt.Sprintf("%d", height*20/scale)

		stylesContent = strings.ReplaceAll(stylesContent, "${x_margin}", x_margin)
		stylesContent = strings.ReplaceAll(stylesContent, "${y_margin}", y_margin)
		stylesContent = strings.ReplaceAll(stylesContent, "${x_hover}", x_hover)
		stylesContent = strings.ReplaceAll(stylesContent, "${y_hover}", y_hover)
	}

	cssPath := filepath.Join(conf.HomeDirPath, ".cache/HyprOne/style.css")
	if err = utils.WriteFile(stylesContent, cssPath); err != nil {
		return err
	}

	command := fmt.Sprintf("wlogout -b %d -c 0 -r 0 -m 0 --layout %s --css %s --protocol layer-shell", cols, layoutPath, cssPath)
	if _, err = utils.ExecCommand(command); err != nil {
		return err
	}

	return nil
}
