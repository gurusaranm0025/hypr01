package init

import (
	battery "gurusaranm0025/hyprone/pkg/modules/Battery"
	"log/slog"
)

func Init() {

	// Battery monitor
	battChannel := make(chan error)

	go battery.BattMon(battChannel)

	for {
		if err := <-battChannel; err != nil {
			slog.Error(err.Error())
		}
	}
}
