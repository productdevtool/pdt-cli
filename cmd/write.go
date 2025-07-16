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

var writeCmd = &cobra.Command{
	Use:   "write [content_type] [topic]",
	Short: "A versatile content generation tool for creating external-facing materials.",
	Long:  "The user specifies a content type (e.g., blog, tweet, landing-page-copy) and a topic.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		contentType := args[0]
		topic := args[1]

		// Build the content generation prompt
		contentPrompt := prompt.ContentGenerationPrompt(contentType, topic)

		color.Cyan("Generating %s content about '%s' with AI...", contentType, topic)
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()

		aiOutput, err := ai.Executor("gemini-cli", contentPrompt)
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

			// Assuming content files go into a 'content' directory
			fullPath := filepath.Join("content", block.FilePath)
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
			color.Green("Wrote content to %s", fullPath)
		}

		color.Green("Content generation complete.")
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
