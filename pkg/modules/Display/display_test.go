package display

import (
	"fmt"
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
	err := Brightness('d')
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
