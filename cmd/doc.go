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

var docCmd = &cobra.Command{
	Use:   "doc [spec_file] [code_paths...]",
	Short: "Instructs the AI to update internal documentation based on a newly implemented feature.",
	Long:  "This command provides the AI with the feature's spec file and the file paths of the implemented code.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		specFile := args[0]
		codePaths := args[1:]

		// Check if spec file exists
		if _, err := os.Stat(specFile); os.IsNotExist(err) {
			color.Red("Error: Spec file '%s' does not exist.", specFile)
			os.Exit(1)
		}

		// Build the doc generation prompt
		docPrompt, err := prompt.DocGenerationPrompt(specFile, codePaths)
		if err != nil {
			color.Red("Error building doc generation prompt: %v", err)
			os.Exit(1)
		}

		color.Cyan("Generating documentation with AI...")
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()

		aiOutput, err := ai.Executor("gemini-cli", docPrompt)
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

			// Assuming documentation files go into a 'docs/handbook' directory
			fullPath := filepath.Join("docs/handbook", block.FilePath)
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
			color.Green("Wrote documentation to %s", fullPath)
		}

		color.Green("Documentation generation complete.")
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}