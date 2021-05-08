package utilities

import (
	"bytes"
	"os/exec"
)

const (
	shell = "bash"
)

func RunConsoleCommand(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(shell, "-c", command)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
