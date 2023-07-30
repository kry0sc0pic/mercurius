
package utils

import (
	"os/exec"
	"errors"
)

// ExecuteCommand executes a shell command and returns an error if it fails
func ExecuteCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Start()
	if err != nil {
		return errors.New("Failed to start command: " + err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		return errors.New("Command execution failed: " + err.Error())
	}

	return nil
}

