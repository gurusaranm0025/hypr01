package setup

import (
	"gurusaranm0025/hyprone/pkg/common"
	"gurusaranm0025/hyprone/pkg/config"
	"gurusaranm0025/hyprone/pkg/modules/themer"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
)

var Dependencies = []string{
	"wireplumber",
	"brightnessctl",
	"swww",
	"cliphist",
	"wl-clipboard",
	"hypridle",
	"eza",
}

func InstallDependencies() error {
	if err := utils.InstallPackages(Dependencies...); err != nil {
		return err
	}
	return nil
}

func DoInitialSetup(force bool) error {
	var err error
	var conf config.Config

	// Initial Setup Check
	if config.CheckInitialSetupNE() && !force {
		slog.Info("Initial Setup is already done. Moving on...")
		return nil
	}

	// CREATE DIRS
	if err = DirsCheck(); err != nil {
		return err
	}

	// INSTALLING DEPENDENCIES
	if err = InstallDependencies(); err != nil {
		return err
	}

	// INSTALL DEFAULT THEME
	theme := themer.NewThemer("default")
	if err = theme.Install(); err != nil {
		return err
	}

	// SAVING CONFIG
	if conf, err = config.LoadConfig(); err != nil {
		return err
	}
	conf.InitialSetup = true
	if err = config.SaveConfig(conf); err != nil {
		return err
	}

	return nil
}

func DirsCheck() error {
	var err error

	for _, path := range common.PlaceholderValues {
		if err = utils.CreateDir(path); err != nil {
			return err
		}
	}
	return nil
}
