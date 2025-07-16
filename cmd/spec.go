package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/ai"
	"github.com/productdevtool/pdt-cli/pkg/prompt"
	"github.com/productdevtool/pdt-cli/pkg/task"
	"github.com/spf13/cobra"
)

var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "Refines a task into a detailed plan using AI.",
	Long:  "This command transforms a high-level todo item into a detailed, actionable technical plan. It focuses on defining the *what* and the *how*",
	Run: func(cmd *cobra.Command, args []string) {
		activeTaskDir, err := task.GetActiveTask()
		if err != nil {
			color.Red("Error getting active task: %v", err)
			os.Exit(1)
		}

		projectDescriptionPath := "docs/project-description.md"
		taskPath := filepath.Join(activeTaskDir, "task.md")

		// Build the refinement prompt
		refinementPrompt, err := prompt.RefineTaskPrompt(projectDescriptionPath, taskPath)
		if err != nil {
			color.Red("Error building refinement prompt: %v", err)
			os.Exit(1)
		}

		color.Cyan("Generating detailed specification with AI...")
		aiOutput, err := ai.Executor("gemini-cli", refinementPrompt)
		if err != nil {
			color.Red("Error executing AI prompt: %v", err)
			os.Exit(1)
		}

		// Update the task.md with the new, detailed plan
		err = ioutil.WriteFile(taskPath, []byte(aiOutput), 0644)
		if err != nil {
			color.Red("Error writing AI output to task.md: %v", err)
			os.Exit(1)
		}

		color.Green("Detailed specification generated and saved to %s", taskPath)

		// Git integration: add and commit the refined task.md
		color.Cyan("Committing refined specification...")
		_, err = ai.Executor("git", "add", taskPath)
		if err != nil {
			color.Red("Error adding task.md to git: %v", err)
			os.Exit(1)
		}

		commitMsg := fmt.Sprintf("feat: Refine spec for %s", filepath.Base(activeTaskDir))
		_, err = ai.Executor("git", "commit", "-m", commitMsg)
		if err != nil {
			color.Red("Error committing task.md: %v", err)
			os.Exit(1)
		}

		color.Green("Refined specification committed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(specCmd)
}
