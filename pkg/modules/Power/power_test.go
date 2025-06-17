package power

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getBatteryPercentAndStatus(t *testing.T) {
	percent, status, err := GetBatteryPercentAndStatus()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Status ==> ", status, "percent ==> ", percent)

	assert.Condition(t, func() bool {
		return statusCheck(status)
	}, fmt.Sprintf("Got \"%s\" instead of Charging, Discharging or Full.", status))

	assert.GreaterOrEqual(t, percent, 0)
	assert.LessOrEqual(t, percent, 100)
}

func statusCheck(status string) bool {
	if status == "Discharging" || status == "Charging" || status == "Full" || status == "Not charging" {
		return true
	}
	return false
}

func Test_BatteryNotifier(t *testing.T) {
	err := BatteryNotifier()
	if err != nil {
		fmt.Println("Error ==> ", err.Error())
	}

	assert.Nil(t, err)
}

func Test_Notify(t *testing.T) {
	err := notify("Battery Low", "58%", true)
	if err != nil {
		fmt.Println(err)
	}

	assert.Nil(t, err)
}
