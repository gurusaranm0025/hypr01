package wallapaper

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/conf"
	"gurusaranm0025/hyprone/pkg/utils"
	"os"
	"path/filepath"
)

// Command for setting the wallpaper
// swww img "$1" --transition-bezier .43,1.19,1,.4 --transition-type "grow" --transition-duration 1 --transition-fps 60 --invert-y --transition-pos "$(hyprctl cursorpos)"

var WALLPAPERS_FOLDER_PATH = filepath.Join(conf.HomeDirPath, "/.HyprOne/Walls")

var WallpaperControls = struct {
	Name             string
	DaemonCMD        string
	SetWall          string
	SetWallAnimFlags string

	GUICommand string
}{
	Name:      "swww",
	DaemonCMD: "swww-daemon",
	// SetWall:    filepath.Join(ScrpitsPath, "setWallpaper.sh"),
	GUICommand: "waypaper --folder " + WALLPAPERS_FOLDER_PATH,
}

func createDirectory(dirPath string) error {
	fmt.Println("creating")
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

func OpenWallGUI() error {
	if _, err := utils.ExecCommand(WallpaperControls.GUICommand, false); err != nil {
		return err
	}
	return nil
}

func StartWallDaemon() error {
	if _, err := utils.ExecCommand(WallpaperControls.DaemonCMD, false); err != nil {
		return err
	}

	if err := createDirectory(WALLPAPERS_FOLDER_PATH); err != nil {
		return err
	}
	return nil
}
