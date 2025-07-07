package display

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBrightness(t *testing.T) {
	percent, err := getBrightness()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Print("Brightness ==> ", percent)

	assert.GreaterOrEqual(t, percent, 0)
	assert.LessOrEqual(t, percent, 100)
}

func Test_setBrightness(t *testing.T) {
	err := setBrightness(50)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)
}

func Test_BrightnessFunc(t *testing.T) {
	err := Brightness("+")
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)
}

func Test_GetScreenRes(t *testing.T) {
	width, height, scale, err := GetScreenresolution()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Width ==> %d, Height ==> %d, Scale ==> %d", width, height, scale)

	assert.Nil(t, err)
}

func Test_GetHyprBorder(t *testing.T) {
	border, err := GetHyprBorder()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Border ==> %d", border)

	assert.Nil(t, err)
}

func Test_BrightnessAll(t *testing.T) {
	var err error

	if err = Brightness("+"); err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	if err = Brightness("5%+"); err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	if err = Brightness("50%"); err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	if err = Brightness("hi+"); err != nil {
		fmt.Println(err.Error())
	}
	assert.NotNil(t, err)

}

func Test_hypridleToggle(t *testing.T) {
	var err error

	process := utils.IsProcessRunning("hypridle")

	if err = ToggleHyprIdle("toggle"); err != nil {
		slog.Error(err.Error())
	}

	assert.NotEqual(t, process, utils.IsProcessRunning("hypridle"), "Toggle hypridle not working.")
}

func Test_hypridleOn(t *testing.T) {
	var err error

	process := utils.IsProcessRunning("hypridle")

	if process {
		if _, err = utils.ExecCommand("pkill hypridle"); err != nil {
			slog.Error(err.Error())
		}
	}

	if err = ToggleHyprIdle("1"); err != nil {
		slog.Error(err.Error())
	}

	assert.Equal(t, true, utils.IsProcessRunning("hypridle"))
}

func Test_hypridleOff(t *testing.T) {
	var err error

	process := utils.IsProcessRunning("hypridle")

	if !process {
		if _, err = utils.ExecCommand("hypridle & disown"); err != nil {
			slog.Error(err.Error())
		}
	}

	if err = ToggleHyprIdle("0"); err != nil {
		slog.Error(err.Error())
	}

	assert.Equal(t, false, utils.IsProcessRunning("hypridle"))
}
