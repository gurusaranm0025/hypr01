package volume

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetVolume(t *testing.T) {
	volumeVal, err := getVolume("speaker")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Volume ==> ", volumeVal)

	assert.IsType(t, 0, volumeVal, "Not a integer.")
	assert.GreaterOrEqual(t, volumeVal, 0)
}

func Test_MuteSpeaker(t *testing.T) {
	if err := Mute("speaker"); err != nil {
		fmt.Println(err.Error())
		fmt.Println("============")
	}

	volumeVal, err := getVolume("speaker")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("============")
	}

	fmt.Println("Volume ==> ", volumeVal)

	assert.IsType(t, 0, volumeVal, "Not an integer.")
	assert.Equal(t, -999, volumeVal, "Device not muted.")
}

func Test_MuteMic(t *testing.T) {
	if err := Mute("mic"); err != nil {
		fmt.Println(err.Error())
		fmt.Println("============")
	}

	volumeVal, err := getVolume("mic")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("============")
	}

	fmt.Println("Volume ==> ", volumeVal)

	assert.IsType(t, 0, volumeVal, "Not an integer.")
	assert.Equal(t, -999, volumeVal, "Device not muted.")
}

func Test_SetVolume(t *testing.T) {
	err := setVolume("speaker", 25)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)
}

func Test_SetVolumeMore100(t *testing.T) {
	err := setVolume("speaker", 102)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)
}

func Test_SetVolumeLess0(t *testing.T) {
	err := setVolume("speaker", -3)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)
}

func Test_VolumeAll(t *testing.T) {
	var err error

	err = Volume("+")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	err = Volume("5%+")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	err = Volume("5%-")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	err = Volume("50%")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

}
