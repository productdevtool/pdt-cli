package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/ai"
	"github.com/productdevtool/pdt-cli/pkg/prompt"
	"github.com/productdevtool/pdt-cli/pkg/task"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Finalizes the work by reviewing, committing, and cleaning up the completed task.",
	Long:  "This command finalizes the work by reviewing, committing, and cleaning up the completed task.",
	Run: func(cmd *cobra.Command, args []string) {
		activeTaskDir, err := task.GetActiveTask()
		if err != nil {
			color.Red("Error getting active task: %v", err)
			os.Exit(1)
		}

		color.Cyan("Displaying git diff for review...")
		gitDiffCmd := exec.Command("git", "diff")
		gitDiffCmd.Stdout = os.Stdout
		gitDiffCmd.Stderr = os.Stderr

		if err := gitDiffCmd.Run(); err != nil {
			color.Red("Error running git diff: %v", err)
			os.Exit(1)
		}

		confirm := false
		promptConfirm := &survey.Confirm{
			Message: color.CyanString("Do you want to commit these changes?"),
		}
		survey.AskOne(promptConfirm, &confirm)

		if !confirm {
			color.Yellow("Commit aborted.")
			return
		}

		color.Cyan("Generating commit message with AI...")
		commitPrompt, err := prompt.CommitMessagePrompt(filepath.Join(activeTaskDir, "task.md"))
		if err != nil {
			color.Red("Error building commit message prompt: %v", err)
			os.Exit(1)
		}

		aiCommitMsg, err := ai.Executor("gemini-cli", commitPrompt)
		if err != nil {
			color.Red("Error generating AI commit message: %v", err)
			os.Exit(1)
		}

		// Clean up AI output for commit message
		aiCommitMsg = strings.TrimSpace(aiCommitMsg)
		if strings.HasPrefix(aiCommitMsg, "```") {
			aiCommitMsg = strings.TrimPrefix(aiCommitMsg, "```")
			aiCommitMsg = strings.TrimSuffix(aiCommitMsg, "```")
			aiCommitMsg = strings.TrimSpace(aiCommitMsg)
		}

		color.Green("Using AI-generated commit message:\n%s", aiCommitMsg)

		// Commit the changes
		gitAddCmd := exec.Command("git", "add", ".")
		gitAddCmd.Stdout = os.Stdout
		gitAddCmd.Stderr = os.Stderr
		if err := gitAddCmd.Run(); err != nil {
			color.Red("Error adding files to git: %v", err)
			os.Exit(1)
		}

		gitCommitCmd := exec.Command("git", "commit", "-m", aiCommitMsg)
		gitCommitCmd.Stdout = os.Stdout
		gitCommitCmd.Stderr = os.Stderr
		if err := gitCommitCmd.Run(); err != nil {
			color.Red("Error committing changes: %v", err)
			os.Exit(1)
		}

		color.Green("Changes committed successfully for task %s.", filepath.Base(activeTaskDir))

		// Implement Task Cleanup (Task 4.4)
		color.Cyan("Cleaning up task workspace...")
		doneDir := "docs/todos/done"
		oldTaskPath := filepath.Join(activeTaskDir, "task.md")
		newTaskPath := filepath.Join(doneDir, filepath.Base(activeTaskDir) + ".md")

		err = os.Rename(oldTaskPath, newTaskPath)
		if err != nil {
			color.Red("Error moving task.md to done directory: %v", err)
			// Don't os.Exit(1) here, as the commit was successful
		}

		err = os.Remove(activeTaskDir)
		if err != nil {
			color.Red("Error removing active task directory: %v", err)
			// Don't os.Exit(1) here, as the commit was successful
		}

		color.Green("Task %s moved to %s and work directory removed.", filepath.Base(activeTaskDir), doneDir)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
