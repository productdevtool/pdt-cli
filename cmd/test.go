package cmd

import (
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/ai"
	"github.com/productdevtool/pdt-cli/pkg/fs"
	"github.com/productdevtool/pdt-cli/pkg/prompt"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [spec_file]",
	Short: "Instructs the AI to write comprehensive tests for a given feature based on its specification file.",
	Long:  "This is useful for generating tests for legacy code that doesn't have a task.md or for adding more tests to an existing feature.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		specFile := args[0]

		// Check if spec file exists
		if _, err := os.Stat(specFile); os.IsNotExist(err) {
			color.Red("Error: Spec file '%s' does not exist.", specFile)
			os.Exit(1)
		}

		// Build the test generation prompt
		testPrompt, err := prompt.TestGenerationPrompt(specFile)
		if err != nil {
			color.Red("Error building test generation prompt: %v", err)
			os.Exit(1)
		}

		color.Cyan("Generating tests with AI...")
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()

		aiOutput, err := ai.Executor("gemini-cli", testPrompt)
		if err != nil {
			s.Stop()
			color.Red("Error executing AI prompt: %v", err)
			os.Exit(1)
		}

		s.Stop()
		color.Green("AI output:\n%s", aiOutput)
		
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

			fullPath := block.FilePath // For tests, assume path is relative to current dir
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
			color.Green("Wrote test code to %s", fullPath)
		}

		color.Green("Test generation complete.")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}