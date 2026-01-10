package audio

import (
	"fmt"
	"gurusaranm0025/hyprone/pkg/utils"
	"path/filepath"
)

func Sinkswitch() error {
	script := filepath.Join(utils.GetHomeDir(), ".local/share/bin/sinkswitch.sh")
	cmd := fmt.Sprintf("kitty --title change-sink -e %s", script)
	if _, err := utils.ExecCommand(cmd); err != nil {
		return err
	}

	return nil
}
