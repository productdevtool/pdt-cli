package prompt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/productdevtool/pdt-cli/pkg/types"
)

// InitialPlannerPrompt creates a prompt to generate a draft spec and clarifying questions.
func InitialPlannerPrompt(userGoal string) (string, error) {
	projectContext, err := getProjectContext()
	if err != nil {
		return "", fmt.Errorf("error getting project context: %w", err)
	}

	prompt := fmt.Sprintf(`
	You are an expert software engineering assistant.
	Your goal is to transform a high-level user request into a detailed technical specification.

	Here is the user's request:
	"%s"

	Here is the current project context (directory structure):
	%s

	Based on this, please perform the following actions:
	1.  **Draft a technical specification:** This should be a markdown document outlining the necessary changes, new files, and overall approach.
	2.  **Ask clarifying questions:** Identify any ambiguities or areas where more information is needed from the user.

	Your response MUST be a JSON object with the following structure:
	{
	  "draftSpec": "<MARKDOWN_CONTENT>",
	  "clarifyingQuestions": [
	    {
	      "questionId": "<UNIQUE_ID>",
	      "question": "<YOUR_QUESTION>"
	    }
	  ]
	}
	`, userGoal, projectContext)

	return prompt, nil
}

func RefineSpecPrompt(userGoal, draftSpec string, answeredQuestions []types.ClarifyingQuestion) (string, error) {
	var answers strings.Builder
	for _, q := range answeredQuestions {
		answers.WriteString(fmt.Sprintf("Q: %s\nA: %s\n\n", q.Question, q.Answer))
	}

	prompt := fmt.Sprintf(`
	You are an expert software engineering assistant.
	You have already provided a draft specification and asked some clarifying questions.
	Now, you have received answers to those questions.

	Original user goal: "%s"

	Initial draft spec:
	%s

	Here are the user's answers to your questions:
	%s

	Please update the draft specification based on these answers.
	The final output should be just the refined markdown specification.
	If you still have questions, please incorporate them into the spec as comments.
	`, userGoal, draftSpec, answers.String())

	return prompt, nil
}

func getProjectContext() (string, error) {
	var context strings.Builder
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && (info.Name() == ".git" || info.Name() == "pdt-dist") {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			context.WriteString(path + "\n")
		}
		return nil
	})
	return context.String(), nil
}

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
	Here is the project description:\n%s\n\n\tHere is the task:\n%s\n\n\tPlease refine the task into a detailed, actionable technical plan. The plan should include specific file locations for code changes, required automated tests, and manual user-facing tests. The output should be a markdown file.
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
	Here is the project description:\n%s\n\n\tHere is the detailed task specification:\n%s\n\n\tPlease implement the task based on the provided project description and detailed specification. \
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

const implementSpecTemplate = `You are an autonomous AI programmer agent with the ability to write to the local filesystem.
Your high-level goal is to implement the feature described in the specification below.

## Specification
{{.SpecContent}}

## Action
1.  Read and understand the entire specification.
2.  Determine all necessary file creations, modifications, and deletions to implement the feature.
3.  Break the work up into small, manageable tasks.
4.  Generate the complete, final code for all required files for each task.
5.  Write the code to the correct file paths. Create directories as needed. Overwrite existing files completely.

## Response
After you have successfully written all files, respond with a summary of the actions you took (e.g., "Created foo.go, Modified bar.go"). Do not output the code you generated in your final response.
`

// ImplementSpecPrompt builds a prompt that instructs an external agent (gemini-cli)
// to implement an entire feature specification.
func ImplementSpecPrompt(specContent string) (string, error) {
	tmpl, err := template.New("implementSpec").Parse(implementSpecTemplate)
	if err != nil {
		return "", err
	}

	data := struct {
		SpecContent string
	}{
		SpecContent: specContent,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
