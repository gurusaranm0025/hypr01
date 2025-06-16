package logout

import (
	"encoding/json"
	"errors"
	"fmt"

	"gurusaranm0025/hyprone/pkg/conf"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func GetScreenResolution() (int, int, float64, error) {
	var out string
	var err error

	cmd := "hyprctl -j monitors"

	if out, err = utils.ExecCommand(cmd); err != nil {
		return -1, -1, -1, err
	}

	var displayJSON []DisplaysJSON

	if err = json.Unmarshal([]byte(out), &displayJSON); err != nil {
		return -1, -1, -1, err
	}

	for _, display := range displayJSON {
		if display.Focused {
			return displayJSON[0].Width, displayJSON[0].Height, displayJSON[0].Scale, nil
		}
	}

	return -1, -1, -1, errors.New("focused display not found")

}

func GetHyprBorder() (int, error) {
	var out string
	var err error

	cmd := "hyprctl -j getoption decoration:rounding"

	if out, err = utils.ExecCommand(cmd); err != nil {
		return -1, err
	}

	var HyprOption HyprOptionJSON

	if err = json.Unmarshal([]byte(out), &HyprOption); err != nil {
		return -1, err
	}

	return HyprOption.Int, nil
}

func createDirectory(dirPath string) error {
	var err error

	_, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func Logout(layout int) error {
	var cols int = 0
	var width, height, hyprborder int
	var scale float64
	var err error

	if width, height, scale, err = GetScreenResolution(); err != nil {
		return err
	}

	scale = scale * 100
	if hyprborder, err = GetHyprBorder(); err != nil {
		return err
	}

	layoutPath := fmt.Sprintf("/home/saran/.config/wlogout/layout_%d", layout)
	stylesPath := fmt.Sprintf("/home/saran/.config/wlogout/style_%d.css", layout)

	cssContent, err := cssFileContent(stylesPath)
	if err != nil {
		return err
	}

	fntSize := fmt.Sprintf("%d", height*2/100)
	active_rad := fmt.Sprintf("%d", hyprborder*5)
	button_rad := fmt.Sprintf("%d", hyprborder*8)

	cssContent = strings.ReplaceAll(cssContent, "${fntSize}", string(fntSize))
	cssContent = strings.ReplaceAll(cssContent, "${active_rad}", string(active_rad))
	cssContent = strings.ReplaceAll(cssContent, "${button_rad}", string(button_rad))
	cssContent = strings.ReplaceAll(cssContent, "$HOME", conf.HomeDirPath)

	switch layout {
	case 1:
		cols = 6
		mgn := fmt.Sprintf("%d", height*28/int(scale))
		hvr := fmt.Sprintf("%d", height*23/int(scale))
		cssContent = strings.ReplaceAll(cssContent, "${mgn}", string(mgn))
		cssContent = strings.ReplaceAll(cssContent, "${hvr}", string(hvr))
	case 2:
		cols = 2
		x_mgn := fmt.Sprintf("%d", width*35/int(scale))
		y_mgn := fmt.Sprintf("%d", height*25/int(scale))
		x_hvr := fmt.Sprintf("%d", width*32/int(scale))
		y_hvr := fmt.Sprintf("%d", height*20/int(scale))

		cssContent = strings.ReplaceAll(cssContent, "${x_mgn}", string(x_mgn))
		cssContent = strings.ReplaceAll(cssContent, "${y_mgn}", string(y_mgn))
		cssContent = strings.ReplaceAll(cssContent, "${x_hvr}", string(x_hvr))
		cssContent = strings.ReplaceAll(cssContent, "${y_hvr}", string(y_hvr))

	}

	cssPath, err := saveCSS(cssContent)
	if err != nil {
		return err
	}

	fmt.Println(layout, layoutPath, cssPath, cols)
	colsStr := fmt.Sprintf("%d", cols)

	logoutCMD := exec.Command("wlogout", "-b", colsStr, "-c", "0", "-r", "0", "-m", "0", "--layout", layoutPath, "--css", cssPath, "--protocol", "layer-shell")
	out, err := logoutCMD.CombinedOutput()

	if err != nil {
		slog.Info("WLOGOUT COMMAND'S OUTPUT ==> ")
		fmt.Println(out)
		fmt.Println("=======================OUTPUT END")
		return err
	}

	return nil
}

func saveCSS(css string) (string, error) {
	var err error

	path := filepath.Join(conf.HomeDirPath, ".cache/HyprOne")
	if err = createDirectory(path); err != nil {
		return "", err
	}

	path = filepath.Join(path, "style.css")
	if err = os.WriteFile(path, []byte(css), os.ModePerm); err != nil {
		return "", err
	}

	return path, nil
}

func cssFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
