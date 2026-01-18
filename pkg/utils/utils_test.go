package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFilesAndDirs(t *testing.T) {
	entries, err := ListFilesAndDirs("/sys/class/sound", false)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}

	assert.IsType(t, []string{}, entries, "Not a string slice")
}

func TestGetFilesAndDirs(t *testing.T) {
	entries, err := GetFilesAndDirs("/sys/class/sound", false)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}

	assert.IsType(t, []Entry{}, entries, "Not a Custom Entry type")
}

func TestPackageInstaller(t *testing.T) {
	err := InstallPackages([]string{"bash"}...)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err, "OH no..., there is an error, it seems...")
}
