package main

import (
	"fmt"
	audio "gurusaranm0025/hyprone/pkg/modules/Audio"
	battery "gurusaranm0025/hyprone/pkg/modules/Battery"
	brightness "gurusaranm0025/hyprone/pkg/modules/Brightness"
	initialise "gurusaranm0025/hyprone/pkg/modules/Init"
	logout "gurusaranm0025/hyprone/pkg/modules/Logout"
	wallapaper "gurusaranm0025/hyprone/pkg/modules/Wallapaper"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var iBright, dBright, iVol, dVol, Init, batStat, screenRes, speakerMuteTog, micMuteTog, wallGUI bool
	var wlogout int

	var rootCMD = &cobra.Command{
		Use:   "hyprone",
		Short: "a package that provides, audio and brightness controls, other small services for WMs.",
		Long:  "My personal tool to control all my devices connected to my PC, made for Hyprland.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if wallGUI {
				if err := wallapaper.OpenWallGUI(); err != nil {
					return err
				}
				return nil
			}

			if micMuteTog {
				if err := audio.MuteMic(); err != nil {
					return err
				}
				return nil
			}

			if speakerMuteTog {
				if err := audio.MuteSpeaker(); err != nil {
					return err
				}
				return nil
			}

			if iVol {
				if err := audio.Volume('i'); err != nil {
					return err
				}
				return nil
			}

			if dVol {
				if err := audio.Volume('d'); err != nil {
					return err
				}
				return nil
			}

			if iBright {
				if err := brightness.Brightness('i'); err != nil {
					return err
				}
				return nil
			}

			if dBright {
				if err := brightness.Brightness('d'); err != nil {
					return err
				}
				return nil
			}

			if wlogout >= 0 {
				if err := logout.Logout(wlogout); err != nil {
					return err
				}
				return nil
			}

			if screenRes {
				width, height, scale, err := logout.GetScreenResolution()
				if err != nil {
					return err
				}
				fmt.Printf("width ==> %d, Height ==> %d in pixels, Scale ==> %f.\n", width, height, scale)
				return nil
			}

			if batStat {
				percent, status, err := battery.GetBatteryChargePercentAndStatus()
				if err != nil {
					return err
				}
				fmt.Printf("Battery Percentage ==> %d, Battery Status ==> %s.\n", percent, status)
				return nil
			}

			if Init {
				initialise.Initialise()
				return nil
			}

			return nil
		},
	}

	rootCMD.Flags().BoolVarP(&iBright, "incr-bright", "B", false, "increases the display brightness.")

	rootCMD.Flags().BoolVarP(&dBright, "decr-bright", "b", false, "dereases the brightness of the display.")

	rootCMD.Flags().BoolVarP(&iVol, "incr-vol", "V", false, "increases the volume of the main speaker.")

	rootCMD.Flags().BoolVarP(&dVol, "decr-vol", "v", false, "decreases the volume of the main speaker.")

	rootCMD.Flags().BoolVarP(&speakerMuteTog, "speaker-toggle", "A", false, "mute and unmute the main audio device")

	rootCMD.Flags().BoolVarP(&micMuteTog, "mic-toggle", "a", false, "mute and unmute the microphone of the main audio device")

	rootCMD.Flags().BoolVarP(&Init, "init", "I", false, "initialise battery monitor and starts wallpaper daemon.")

	rootCMD.Flags().BoolVarP(&batStat, "battery-stat", "s", false, "tells the battery charge level and its status")

	rootCMD.Flags().BoolVarP(&screenRes, "display-resolution", "D", false, "prints the display's resolution")

	rootCMD.Flags().BoolVarP(&wallGUI, "wall-gui", "W", false, "open's wallpaper changing GUI - waypaper")

	rootCMD.Flags().IntVarP(&wlogout, "logout", "L", -1, "opens logout dialog.")

	if err := rootCMD.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
		return
	}

	os.Exit(0)
}
