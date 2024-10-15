package Init

import (
	battery "gurusaranm0025/hyprone/pkg/modules/Battery"
	wallapaper "gurusaranm0025/hyprone/pkg/modules/Wallapaper"
	"log/slog"
)

func Init() {
	if err := wallapaper.StartWallDaemon(); err != nil {
		slog.Error(err.Error())
	}

	// Battery monitor
	battChannel := make(chan error)

	go battery.BattMon(battChannel)

	for {
		if err := <-battChannel; err != nil {
			slog.Error(err.Error())
		}
	}
}
