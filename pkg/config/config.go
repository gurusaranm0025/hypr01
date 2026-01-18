package config

import (
	"encoding/json"
	"errors"
	"gurusaranm0025/hyprone/pkg/common"
	"log/slog"
	"os"
)

var IsADir = errors.New("provided is a directory")

type Config struct {
	InitialSetup bool `json:"initial-setup"`
}

func LoadConfig() (Config, error) {
	var config Config

	data, err := os.ReadFile(common.HYPR01_CONFIG_PATH)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func SaveConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(common.HYPR01_CONFIG_PATH, data, 0644)
}

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

	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return false, err
	}

	return conf.InitialSetup, nil
}

func CheckInitialSetupNE() bool {
	isDone, err := CheckInitialSetup()
	if err != nil {
		slog.Error(err.Error())
	}

	return isDone
}
