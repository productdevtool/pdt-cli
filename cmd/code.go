package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/ai"
	"github.com/productdevtool/pdt-cli/pkg/fs"
	"github.com/productdevtool/pdt-cli/pkg/prompt"
	"github.com/productdevtool/pdt-cli/pkg/task"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Executes the implementation plan, instructing the AI to perform the necessary coding, testing, and validation.",
	Long:  "This command executes the implementation plan, instructing the AI to perform the necessary coding, testing, and validation.",
	Run: func(cmd *cobra.Command, args []string) {
		activeTaskDir, err := task.GetActiveTask()
		if err != nil {
			color.Red("Error getting active task: %v", err)
			os.Exit(1)
		}

		projectDescriptionPath := "docs/project-description.md"
		taskPath := filepath.Join(activeTaskDir, "task.md")

		// Build the master implementation prompt
		masterPrompt, err := prompt.MasterImplementationPrompt(projectDescriptionPath, taskPath)
		if err != nil {
			color.Red("Error building master implementation prompt: %v", err)
			os.Exit(1)
		}

		color.Cyan("Generating code with AI...")
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()

		aiOutput, err := ai.Executor("gemini-cli", masterPrompt)
		if err != nil {
			s.Stop()
			color.Red("Error executing AI prompt: %v", err)
			os.Exit(1)
		}

		s.Stop()
		color.Green("AI output received. Parsing code blocks...")

		codeBlocks, err := fs.ExtractCodeBlocks(aiOutput)
		if err != nil {
			color.Red("Error extracting code blocks from AI output: %v", err)
			os.Exit(1)
		}

		for _, block := range codeBlocks {
			if block.FilePath == "" {
				color.Yellow("Skipping code block with no file path:\n%s", block.Content)
				continue
			}

			fullPath := filepath.Join(activeTaskDir, block.FilePath)
			// Ensure directory exists
			err = os.MkdirAll(filepath.Dir(fullPath), 0755)
			if err != nil {
				color.Red("Error creating directory for %s: %v", fullPath, err)
				continue
			}

			// Write content to file
			err = os.WriteFile(fullPath, []byte(block.Content), 0644)
			if err != nil {
				color.Red("Error writing to file %s: %v", fullPath, err)
				continue
			}
			color.Green("Wrote code to %s", fullPath)
		}

		color.Green("Code generation complete.")

		// Automated Validation (Task 4.2)
		validationCommands, err := fs.GetValidationCommands()
		if err != nil {
			color.Red("Error getting validation commands: %v", err)
			os.Exit(1)
		}

		if len(validationCommands) > 0 {
			color.Cyan("Running automated validation...")
			for _, valCmd := range validationCommands {
				color.Cyan("Executing: %s", valCmd)
				cmdParts := strings.Fields(valCmd)
				valExecCmd := exec.Command(cmdParts[0], cmdParts[1:]...)
				valExecCmd.Stdout = os.Stdout
				valExecCmd.Stderr = os.Stderr

				if err := valExecCmd.Run(); err != nil {
					color.Red("Validation command failed: %v", err)
					// TODO: Implement AI re-prompting and retry logic here
					color.Yellow("Automated validation failed. Please review the output and fix the issues.")
					os.Exit(1)
				}
			}
			color.Green("Automated validation passed.")
		} else {
			color.Yellow("No automated validation commands found in project-description.md.")
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
}
