package utils

import (
	"os/exec"
	"strings"
)

// function to execute shell commands
func ExecCommand(command string, wantOutput bool) (string, error) {
	var cmdOutput []byte
	var err error

	commandSplits := strings.Split(command, " ")

	cmd := exec.Command(commandSplits[0], commandSplits[1:]...)

	if wantOutput {
		cmdOutput, err = cmd.CombinedOutput()
	} else {
		err = cmd.Start()
	}

	if err != nil {
		return "", err
	}

	return string(cmdOutput), nil

}
