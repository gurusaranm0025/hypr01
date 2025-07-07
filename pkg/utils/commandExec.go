package utils

import (
	"os/exec"
	"strings"
)

func ExecCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func ExecInBackground(command string) error {
	cmd := exec.Command("bash", "-c", strings.TrimSpace(command)+" & disown")
	err := cmd.Run()
	return err
}
