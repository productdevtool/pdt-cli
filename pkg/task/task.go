package task

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// GetActiveTask returns the path to the active task directory.
// It returns an error if there is not exactly one task in the work directory.
func GetActiveTask() (string, error) {
	workDir := "docs/todos/work"
	files, err := ioutil.ReadDir(workDir)
	if err != nil {
		return "", err
	}

	var taskDirs []string
	for _, file := range files {
		if file.IsDir() {
			taskDirs = append(taskDirs, file.Name())
		}
	}

	if len(taskDirs) == 0 {
		return "", fmt.Errorf("no active task found in %s", workDir)
	}

	if len(taskDirs) > 1 {
		return "", fmt.Errorf("multiple active tasks found in %s, please specify which one to use", workDir)
	}

	return filepath.Join(workDir, taskDirs[0]), nil
}
