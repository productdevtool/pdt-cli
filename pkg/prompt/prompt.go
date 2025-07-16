package prompt

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// RefineTaskPrompt generates a prompt for refining a task.md file.
func RefineTaskPrompt(projectDescriptionPath string, taskPath string) (string, error) {
	projectDescription, err := ioutil.ReadFile(projectDescriptionPath)
	if err != nil {
		return "", fmt.Errorf("error reading project description: %w", err)
	}

	task, err := ioutil.ReadFile(taskPath)
	if err != nil {
		return "", fmt.Errorf("error reading task: %w", err)
	}

	prompt := fmt.Sprintf(`\
	Here is the project description:\n%s\n\n	Here is the task:\n%s\n\n	Please refine the task into a detailed, actionable technical plan. The plan should include specific file locations for code changes, required automated tests, and manual user-facing tests. The output should be a markdown file.
	`, string(projectDescription), string(task))

	return prompt, nil
}

// InitialProjectDescriptionPrompt generates a prompt for creating the initial project-description.md.
func InitialProjectDescriptionPrompt() string {
	return `\
	Please analyze the current codebase and generate a comprehensive project description. \
	This description should include the project's structure, key commands, main functionalities, \
	and any other relevant information that would help an AI understand and work with this project. \
	The output should be a markdown file named "project-description.md".
	`
}

// MasterImplementationPrompt generates the master prompt for code generation.
func MasterImplementationPrompt(projectDescriptionPath string, taskPath string) (string, error) {
	projectDescription, err := ioutil.ReadFile(projectDescriptionPath)
	if err != nil {
		return "", fmt.Errorf("error reading project description: %w", err)
	}

	task, err := ioutil.ReadFile(taskPath)
	if err != nil {
		return "", fmt.Errorf("error reading task: %w", err)
	}

	prompt := fmt.Sprintf(`\
	Here is the project description:\n%s\n\n	Here is the detailed task specification:\n%s\n\n	Please implement the task based on the provided project description and detailed specification. \
	Generate the necessary code, making sure to adhere to the specified file locations and include any required tests. \
	Provide the output as code blocks, clearly indicating file paths for each code block.
	`, string(projectDescription), string(task))

	return prompt, nil
}

// CommitMessagePrompt generates a prompt for creating a commit message.
func CommitMessagePrompt(taskPath string) (string, error) {
	task, err := ioutil.ReadFile(taskPath)
	if err != nil {
		return "", fmt.Errorf("error reading task: %w", err)
	}

	prompt := fmt.Sprintf(`\
	Based on the following task specification, please generate a concise and descriptive Git commit message. \
	Focus on the "what" and "why" of the changes. The commit message should follow conventional commits guidelines (e.g., feat: add new feature). \
	Task: %s
	`, string(task))

	return prompt, nil
}

// TestGenerationPrompt generates a prompt for creating tests based on a spec file.
func TestGenerationPrompt(specPath string) (string, error) {
	specContent, err := ioutil.ReadFile(specPath)
	if err != nil {
		return "", fmt.Errorf("error reading spec file: %w", err)
	}

	prompt := fmt.Sprintf(`\
	Based on the following specification, please generate comprehensive tests. \
	The tests should cover unit, integration, and end-to-end scenarios as appropriate. \
	Adhere to the project's existing testing patterns and frameworks. \
	Provide the output as code blocks, clearly indicating file paths for each test file.
	Specification: %s
	`, string(specContent))

	return prompt, nil
}

// DocGenerationPrompt generates a prompt for updating internal documentation.
func DocGenerationPrompt(specPath string, codePaths []string) (string, error) {
	specContent, err := ioutil.ReadFile(specPath)
	if err != nil {
		return "", fmt.Errorf("error reading spec file: %w", err)
	}

	var codeContents []string
	for _, path := range codePaths {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("error reading code file %s: %w", path, err)
		}
		codeContents = append(codeContents, fmt.Sprintf("File: %s\n```\n%s\n```", path, string(content)))
	}

	prompt := fmt.Sprintf(`\
	Based on the following specification and implemented code, please update internal documentation. \
	Explain how the feature works, its API, and how to use it. \
	Specification: %s\n\nImplemented Code:\n%s
	`, string(specContent), strings.Join(codeContents, "\n\n"))

	return prompt, nil
}

// ContentGenerationPrompt generates a prompt for creating external-facing content.
func ContentGenerationPrompt(contentType string, topic string) string {
	return fmt.Sprintf(`\
	Generate %s content about the following topic: %s. \
	The output should be suitable for direct use and saved to a new file in a /content directory. \
	`, contentType, topic)
}
