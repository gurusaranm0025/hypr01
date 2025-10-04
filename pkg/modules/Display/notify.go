package display

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"log/slog"
)

func notify() error {

	percent, err := getBrightness()
	if err != nil {
		return err
	}

	command := fmt.Sprintf("swayosd-client --custom-icon=display-brightness --custom-progress=%.2f", max(float32(percent)/float32(100), 0.01))

	if output, err := utils.ExecCommand(command); err != nil {
		slog.Error(output)
		return err
	}

	return nil
}
