package utils

import (
	"log/slog"
	"os/exec"
)

func InstallPackages(pkgs ...string) error {
	args := append([]string{"pacman", "-S", "--needed", "--noconfirm"}, pkgs...)

	cmd := exec.Command("pkexec", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(string(out))
		return err
	}

	return nil
}
