package utils

import (
	"os/exec"
)

func Execute(path string) (int, error) {
	cmd := exec.Command(path)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), err
		}
		return -1, err
	}
	return 0, nil
}
