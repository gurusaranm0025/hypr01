package power

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

var LOWPOINTS = []int{1, 2, 3, 4, 5, 10, 15, 17, 20, 25}
var HIGHPOINTS = []int{80, 85, 90, 95, 99}
var lastNotificationPercent = 0

var BATTERYMODES = struct {
	Charging    string
	Discharging string
	Full        string
	NotCharging string
}{
	Charging:    "Charging",
	Discharging: "Discharging",
	Full:        "Full",
	NotCharging: "Not charging",
}

func GetBatteryPercentAndStatus() (int, string, error) {
	var percent int
	var status string

	path := "/sys/class/power_supply"
	entries, err := utils.GetFilesAndDirs(path, false)
	if err != nil {
		return -1, "", nil
	}

	var filePath string
	var file []byte
	for _, entry := range entries {
		if strings.HasPrefix(entry.Info.Name(), "BAT") && !entry.Info.IsDir() {
			// Battery charge percentage
			filePath = filepath.Join(entry.Path, "capacity")
			if file, err = os.ReadFile(filePath); err != nil {
				return -1, "", err
			}

			if percent, err = strconv.Atoi(strings.TrimSpace(string(file))); err != nil {
				return -1, "", err
			}

			// Battery Status
			filePath = filepath.Join(entry.Path, "status")
			if file, err = os.ReadFile(filePath); err != nil {
				return -1, "", nil
			}

			status = strings.TrimSpace(string(file))

			break
		}
	}

	return percent, status, nil
}

func BatteryNotifier() error {
	var err error

	percent, status, err := GetBatteryPercentAndStatus()
	if err != nil {
		return err
	}

	if slices.Contains(LOWPOINTS, percent) && status != BATTERYMODES.Charging && lastNotificationPercent != percent {
		lastNotificationPercent = percent
		err = notify("Battery Low", fmt.Sprintf("%d%% is low. Connect to charger", percent), true)
	} else if slices.Contains(HIGHPOINTS, percent) && status == BATTERYMODES.Charging && lastNotificationPercent != percent {
		lastNotificationPercent = percent
		err = notify("Charging", strconv.Itoa(percent), false)
	} else if (status == BATTERYMODES.Charging || status == BATTERYMODES.NotCharging) && lastNotificationPercent != percent {
		lastNotificationPercent = percent
		err = notify("Fully Charged", "You can unplug the charger", false)
		time.Sleep(10 * time.Minute)
	}

	if err != nil {
		return err
	}

	return nil
}

func BatteryMonitor(ch chan<- error) {
	for {
		if err := BatteryNotifier(); err != nil {
			ch <- err
		}
		time.Sleep(1 * time.Minute)
	}
}
