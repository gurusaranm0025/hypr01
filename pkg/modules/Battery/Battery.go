package battery

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

var LOWPOINTS = []int{1, 2, 3, 4, 5, 10, 12, 15, 17, 20, 25}
var HIGHPOINTS = []int{80, 85, 90, 93, 95, 97, 99}
var chargeAtLastNoti = 0

var BatteryModes = struct {
	Charging    string
	Discharging string
	Full        string
}{
	Charging:    "Charging",
	Discharging: "Discharging",
	Full:        "Full",
}

// function to get the charge percentage and battery charging status
func getBatteryChargePercentAndStatus() (int, string, error) {
	var outPercent int64 = -1
	var outStatus string = ""

	// directory path for battery details
	dirPath := "/sys/class/power_supply/"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return -1, "", err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "BAT") {
			// PERCENT
			percentPath := filepath.Join(dirPath, file.Name(), "capacity")

			percentFile, err := os.ReadFile(percentPath)
			if err != nil {
				return -1, "", err
			}

			outPercent, err = strconv.ParseInt(strings.TrimSpace(string(percentFile)), 10, 0)

			// STATUS
			statusPath := filepath.Join(dirPath, file.Name(), "status")

			statusFile, err := os.ReadFile(statusPath)
			if err != nil {
				return -1, "", err
			}

			outStatus = strings.TrimSpace(string(statusFile))

			break
		}
	}

	return int(outPercent), outStatus, nil
}

// funciton to send a notification
func notify(title, message string, isCritical bool) error {

	var cmd *exec.Cmd

	if isCritical {
		cmd = exec.Command("notify-send", "-t", "5000", "-r", "69", "-u", "CRITICAL", title, message)
	} else {
		cmd = exec.Command("notify-send", "-t", "5000", "-r", "69", title, message)

	}

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

// function to check the battery level and notify the user
func BatteryCheck() error {
	var err error

	battPercent, battStatus, err := getBatteryChargePercentAndStatus()
	if err != nil {
		return err
	}

	if slices.Contains(LOWPOINTS, battPercent) && battStatus != BatteryModes.Charging && chargeAtLastNoti != battPercent {
		chargeAtLastNoti = battPercent
		err = notify("Battery Low", fmt.Sprintf("%d%% is critically low. Connect charger.", battPercent), true)
	}

	if slices.Contains(HIGHPOINTS, battPercent) && battStatus == BatteryModes.Charging && chargeAtLastNoti != battPercent {
		chargeAtLastNoti = battPercent
		err = notify("Battery Charging", fmt.Sprintf("Battery is at %d%% and is still charging.", battPercent), false)
	}

	if battStatus == BatteryModes.Full && chargeAtLastNoti != battPercent {
		chargeAtLastNoti = battPercent
		err = notify("Fully Charged", "Battery is fully charged. You can unplug the charger.", false)
		time.Sleep(10 * time.Minute)
	}

	if err != nil {
		return err
	}

	return nil
}

func BattMon(ch chan<- error) {
	for {
		err := BatteryCheck()
		if err != nil {
			ch <- err
		}
		time.Sleep(45 * time.Second)
	}
}
