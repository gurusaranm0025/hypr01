package cmd

import (
	"fmt"
	display "gurusaranm0025/hyprone/pkg/modules/Display"
	initialize "gurusaranm0025/hyprone/pkg/modules/Initialize"
	logout "gurusaranm0025/hyprone/pkg/modules/Logout"
	volume "gurusaranm0025/hyprone/pkg/modules/Volume"
	wallapaper "gurusaranm0025/hyprone/pkg/modules/Wallapaper"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var VERSION = "0.5.3-4"

var brightness, sound, mute, hypridle string
var initialise, wallpaperGUI, ver bool
var logoutLayout int

var rootCMD = &cobra.Command{
	Use:   "hyprone",
	Short: "aggregator of volume, brightness, wallpaper, battery monitor and logout commands for HyprLand",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if len(brightness) > 0 {
			if err = display.Brightness(brightness); err != nil {
				return err
			}
		}

		if hypridle != "" {
			if err = display.ToggleHyprIdle(hypridle); err != nil {
				return err
			}
		}

		if len(sound) > 0 {
			if err = volume.Volume(sound); err != nil {
				return err
			}
		}

		if len(mute) > 0 {
			if err = volume.Mute(mute); err != nil {
				return err
			}
		}

		if logoutLayout > 0 {
			if err = logout.Logout(logoutLayout); err != nil {
				return err
			}
		}

		if initialise {
			initialize.Initializer()
		}

		if wallpaperGUI {
			if err = wallapaper.WallpaperGUI(); err != nil {
				return err
			}
		}

		if ver {
			fmt.Println(VERSION)
		}

		return nil
	},
}

func initializeFlags() {
	rootCMD.Flags().StringVarP(&brightness, "brightness", "b", "", "Example values: +, -, 5%+, 5%-, 50%")

	rootCMD.Flags().StringVar(&hypridle, "hypridle", "", "Accepted values: toggle - for toggle, 0 - for stopping hypridle, 1 - for running hypridle")

	rootCMD.Flags().StringVarP(&sound, "volume", "v", "", "Example values: +, -, 5%+, 5%-, 50%")

	rootCMD.Flags().StringVarP(&mute, "mute-toggle", "m", "", "values: speaker, mic")

	rootCMD.Flags().BoolVarP(&initialise, "initialize", "i", false, "Initialises Battery Monitor and starts swww-daemon")

	rootCMD.Flags().IntVarP(&logoutLayout, "logout-menu", "l", 0, "values: 1, 2")

	rootCMD.Flags().BoolVarP(&wallpaperGUI, "wallpaper-app", "w", false, "Opens waypaper - wallpaper choosing app")

	rootCMD.Flags().BoolVar(&ver, "version", false, "current version")
}

func Execute() {
	initializeFlags()
	if err := rootCMD.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
