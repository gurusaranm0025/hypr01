package utils

import (
	"fmt"
	"strings"
)

func IsProcessRunning(name string) bool {
	command := fmt.Sprintf("pgrep -x %s", name)
	out, err := ExecCommand(command)

	return (err == nil && strings.TrimSpace(string(out)) != "")
}
