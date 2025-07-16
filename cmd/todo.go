package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/ai"
	"github.com/productdevtool/pdt-cli/pkg/fs"
	"github.com/productdevtool/pdt-cli/pkg/prompt"
	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Starts the workflow, generates initial project context, and allows task selection.",
	Long:  `This is the starting point for any new work. It initializes the environment, allows the user to select a task, and prepares the workspace.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initWorkspace(); err != nil {
			color.Red("Error initializing workspace: %v", err)
			os.Exit(1)
		}

		tasks, err := fs.ReadTodoFile("docs/todo.md")
		if err != nil {
			color.Red("Error reading todo file: %v", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			color.Yellow("No tasks found in docs/todo.md. Add some tasks and try again.")
			return
		}

		var selectedTask string
		prompt := &survey.Select{
			Message: color.CyanString("Choose a task to begin:"),
			Options: tasks,
		}
		survey.AskOne(prompt, &selectedTask)

		if err := initializeTaskWorkspace(selectedTask, tasks); err != nil {
			color.Red("Error initializing task workspace: %v", err)
			os.Exit(1)
		}

		color.Green("Successfully initialized workspace for task: %s", selectedTask)
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}

func initWorkspace() error {
	if err := fs.CreateDirs([]string{"docs", "docs/todos/work", "docs/todos/done"}); err != nil {
		return err
	}

	// Check for project-description.md and generate if not found
	projectDescExists, err := fs.Exists("docs/project-description.md")
	if err != nil {
		return err
	}

	if !projectDescExists {
		color.Cyan("Generating initial project description...")
		initialPrompt := prompt.InitialProjectDescriptionPrompt()
		aiOutput, err := ai.Executor("gemini-cli", initialPrompt)
		if err != nil {
			return fmt.Errorf("failed to generate project description: %w", err)
		}

		// Write the AI output to the file
		err = os.WriteFile("docs/project-description.md", []byte(aiOutput), 0644)
		if err != nil {
			return fmt.Errorf("failed to write project description to file: %w", err)
		}

		color.Green("Initial project description generated in docs/project-description.md.")

		confirmEdit := false
		promptConfirmEdit := &survey.Confirm{
			Message: color.CyanString("Do you want to review and edit the generated project-description.md?"),
		}
		survey.AskOne(promptConfirmEdit, &confirmEdit)

		if confirmEdit {
			// TODO: Implement logic to open editor for project-description.md
			color.Yellow("Please manually edit docs/project-description.md and then run pdt todo again.")
			os.Exit(0)
		}
	}

	// Ensure docs/todo.md exists
	todoExists, err := fs.Exists("docs/todo.md")
	if err != nil {
		return err
	}

	if !todoExists {
		_, err := os.Create("docs/todo.md")
		if err != nil {
			return err
		}
	}

	return nil
}

func initializeTaskWorkspace(task string, allTasks []string) error {
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	taskName := strings.ToLower(strings.ReplaceAll(task, " ", "-"))
	taskDir := fmt.Sprintf("docs/todos/work/%s-%s", timestamp, taskName)

	if err := os.MkdirAll(taskDir, 0755); err != nil {
		return err
	}

	taskFile, err := os.Create(fmt.Sprintf("%s/task.md", taskDir))
	if err != nil {
		return err
	}
	defer taskFile.Close()

	_, err = taskFile.WriteString(fmt.Sprintf("# Task: %s\n", task))
	if err != nil {
		return err
	}

	var remainingTasks []string
	for _, t := range allTasks {
		if t != task {
			remainingTasks = append(remainingTasks, t)
		}
	}

	if err := fs.RewriteTodoFile("docs/todo.md", remainingTasks); err != nil {
		return err
	}

	return nil
}