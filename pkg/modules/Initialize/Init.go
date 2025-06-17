package initialize

import (
	power "gurusaranm0025/hyprone/pkg/modules/Power"
	wallapaper "gurusaranm0025/hyprone/pkg/modules/Wallapaper"
	"log/slog"
)

func Initializer() {
	if err := wallapaper.StartDaemon(); err != nil {
		slog.Error(err.Error())
	}

	// Battery monitor
	battChannel := make(chan error)

	go power.BatteryMonitor(battChannel)

	for {
		if err := <-battChannel; err != nil {
			slog.Error(err.Error())
		}
	}
}
