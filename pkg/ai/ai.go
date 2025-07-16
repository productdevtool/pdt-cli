package ai

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// Executor executes a command and returns its stdout as a string.
// It also streams stdout and stderr to the console.
func Executor(command string, args ...string) (string, error) {
	var cmd *exec.Cmd

	// Get the directory of the currently running executable
	ex, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}
	exDir := filepath.Dir(ex)

	// Construct the path to the command in the same directory
	localCmdPath := filepath.Join(exDir, command)

	// Check if the command exists at the local path
	if _, err := os.Stat(localCmdPath); err == nil {
		// If it exists, use this path
		cmd = exec.Command(localCmdPath, args...)
	} else {
		// Otherwise, look for the command in the system's PATH
		cmd = exec.Command(command, args...)
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("error starting command '%s': %w", command, err)
	}

	if err := cmd.Wait(); err != nil {
		// For gemini-cli, consider parsing stderr for more specific error messages
		return "", fmt.Errorf("command '%s' finished with error: %w\nStderr: %s", command, err, stderrBuf.String())
	}

	// Further refinement for gemini-cli output might involve parsing JSON or other structured data
	// based on its specific output formats. For now, we return the raw string.
	return stdoutBuf.String(), nil
}

