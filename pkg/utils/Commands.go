package utils

import (
	"os"
	"os/exec"
)

// function to execute shell commands
// func ExecCommand(command string, wantOutput bool) (string, error) {
// 	var cmdOutput []byte
// 	var err error
// 	var cmd *exec.Cmd

// 	commandSplits := strings.Split(command, " ")
// 	cmd = exec.Command(commandSplits[0], commandSplits[1:]...)

// 	if wantOutput {
// 		cmdOutput, err = cmd.CombinedOutput()
// 	} else {
// 		err = cmd.Start()
// 	}

// 	if err != nil {
// 		return "", err
// 	}

// 	return string(cmdOutput), nil

// }

func ExecCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// function to execute sudo commands
func SudoCommand(command string) error {

	// commandSplits := strings.Split(command, " ")
	// fmt.Println("CommandSplits ==>")
	// fmt.Println(commandSplits)
	cmd := exec.Command("bash", "-c", command)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
