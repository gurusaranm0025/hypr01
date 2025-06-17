package power

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
)

func notify(title, msg string, isCritical bool) error {
	var command string

	if isCritical {
		command = fmt.Sprintf("notify-send -a \"hyprone\" -t 5000 -r 000003 -u CRITICAL \"%s\" \"%s\"", title, msg)
	} else {
		command = fmt.Sprintf("notify-send -t 5000 -r 000003 \"%s\" \"%s\"", title, msg)
	}

	if _, err := utils.ExecCommand(command); err != nil {
		return err
	}

	return nil
}
