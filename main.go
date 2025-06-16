package main

import (
	"errors"
	"fmt"
	battery "gurusaranm0025/hyprone/pkg/modules/Battery"
	brightness "gurusaranm0025/hyprone/pkg/modules/Brightness"
	Init "gurusaranm0025/hyprone/pkg/modules/Init"
	logout "gurusaranm0025/hyprone/pkg/modules/Logout"
	volume "gurusaranm0025/hyprone/pkg/modules/Volume"
	wallapaper "gurusaranm0025/hyprone/pkg/modules/Wallapaper"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var iBright, dBright, iVol, dVol, initialise, batStat, screenRes, speakerMuteTog, micMuteTog, wallGUI, shell_hist, version bool
	var wlogout int
	var powerProfileMode string

	var rootCMD = &cobra.Command{
		Use:   "hyprone",
		Short: "a package that provides, volume and brightness controls, other small services for WMs.",
		Long:  "My personal tool to control all my devices connected to my PC, made for Hyprland.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if wallGUI {
				if err := wallapaper.OpenWallGUI(); err != nil {
					return err
				}
			}

			if micMuteTog {
				if err := volume.Mute("mic"); err != nil {
					return err
				}
			}

			if speakerMuteTog {
				if err := volume.Mute("speaker"); err != nil {
					return err
				}
			}

			if iVol {
				if err := volume.Volume('i'); err != nil {
					return err
				}
			}

			if dVol {
				if err := volume.Volume('d'); err != nil {
					return err
				}
			}

			if iBright {
				if err := brightness.Brightness('i'); err != nil {
					return err
				}
			}

			if dBright {
				if err := brightness.Brightness('d'); err != nil {
					return err
				}
			}

			if wlogout >= 0 {
				if err := logout.Logout(wlogout); err != nil {
					return err
				}
			}

			if screenRes {
				width, height, scale, err := logout.GetScreenResolution()
				if err != nil {
					return err
				}
				fmt.Printf("width ==> %d, Height ==> %d in pixels, Scale ==> %f.\n", width, height, scale)
			}

			if batStat {
				percent, status, err := battery.GetBatteryChargePercentAndStatus()
				if err != nil {
					return err
				}
				fmt.Printf("Battery Percentage ==> %d, Battery Status ==> %s.\n", percent, status)
			}

			if shell_hist {
				_, err := utils.ExecCommand("bash -c /opt/hyprone/Scripts/zsh_shell_history.sh")
				if err != nil {
					return err
				}
			}

			if version {
				fmt.Println("0.4.3")
			}

			if initialise {
				fmt.Print(121212122121)
				Init.Init()
			}

			if len(powerProfileMode) > 0 {
				// err := performace.ChangePowerProfile(powerProfileMode)
				// if err != nil {
				// 	return err
				// }
				return errors.New("power profiles not available")
			}

			return nil
		},
	}

	rootCMD.Flags().BoolVarP(&iBright, "incr-bright", "B", false, "increases the display brightness.")

	rootCMD.Flags().BoolVarP(&dBright, "decr-bright", "b", false, "dereases the brightness of the display.")

	rootCMD.Flags().BoolVarP(&iVol, "incr-vol", "V", false, "increases the volume of the main speaker.")

	rootCMD.Flags().BoolVarP(&dVol, "decr-vol", "v", false, "decreases the volume of the main speaker.")

	rootCMD.Flags().BoolVarP(&speakerMuteTog, "speaker-toggle", "A", false, "mute and unmute the main volume device")

	rootCMD.Flags().BoolVarP(&micMuteTog, "mic-toggle", "a", false, "mute and unmute the microphone of the main volume device")

	rootCMD.Flags().BoolVarP(&initialise, "init", "I", false, "initialise battery monitor and starts wallpaper daemon.")

	rootCMD.Flags().BoolVarP(&batStat, "battery-stat", "s", false, "tells the battery charge level and its status")

	rootCMD.Flags().BoolVarP(&screenRes, "display-resolution", "D", false, "prints the display's resolution")

	rootCMD.Flags().BoolVarP(&wallGUI, "wall-gui", "W", false, "open's wallpaper changing GUI - waypaper")

	rootCMD.Flags().BoolVarP(&shell_hist, "shell-history", "", false, "shows zsh shell history, in rofi")

	rootCMD.Flags().BoolVarP(&version, "version", "", false, "version of the package")

	rootCMD.Flags().IntVarP(&wlogout, "logout", "L", -1, "opens logout dialog.")

	rootCMD.Flags().StringVarP(&powerProfileMode, "power-profile", "p", "", "change power profile (power saving, balanced, performance)")

	if err := rootCMD.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
		return
	}

	os.Exit(0)
}
