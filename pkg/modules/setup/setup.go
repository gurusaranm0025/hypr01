package setup

import (
	"encoding/json"
	"errors"
	"gurusaranm0025/hyprone/pkg/common"
	"gurusaranm0025/hyprone/pkg/config"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"os"
)

var IsADir = errors.New("provided is a directory")

func CheckInitialSetup() (bool, error) {
	info, err := os.Stat(common.HYPR01_CONFIG_PATH)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if info.IsDir() {
		return false, IsADir
	}

	// Reading the config file
	data, err := os.ReadFile(common.HYPR01_CONFIG_PATH)
	if err != nil {
		return false, err
	}

	var cfg config.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return false, err
	}

	if cfg.InitialSetup {
		return true, nil
	}

	return false, nil
}

func CheckInitialSetupNE() bool {
	isDone, err := CheckInitialSetup()
	if err != nil {
		slog.Error(err.Error())
	}

	return isDone
}

func DoInitialSetup() {
	// CREATE DIRS
	utils.CreateDir(common.CONFIG_DIR_PATH)
	utils.CreateDir(common.ALL_WALLS_DIR_PATH)
	utils.CreateDir(common.CURRENT_WALL_DIR_PATH)

}
