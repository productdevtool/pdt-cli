package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/ai"
	"github.com/productdevtool/pdt-cli/pkg/prompt"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code [spec_file_path]",
	Short: "Implements a feature specification using an AI agent.",
	Long:  "This command reads a high-level feature specification and instructs an AI agent (gemini cli) to generate and write all necessary code to the filesystem.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Reminder: For safety, it's recommended to run this command on a dedicated feature branch.")

		specPath := ".pdt/specs/spec.md" // Default spec path
		if len(args) > 0 {
			specPath = args[0]
		}

		color.Cyan("Reading specification from: %s", specPath)

		specContent, err := os.ReadFile(specPath)
		if err != nil {
			color.Red("Error reading spec file at %s: %v", specPath, err)
			color.Yellow("Hint: Ensure the spec file exists. You can generate one with `pdt spec`.")
			os.Exit(1)
		}

		if len(specContent) == 0 {
			color.Yellow("Spec file is empty. Nothing to do.")
			os.Exit(0)
		}

		// Build the prompt that passes the entire spec to the agent.
		codePrompt, err := prompt.ImplementSpecPrompt(string(specContent))
		if err != nil {
			color.Red("Error building prompt: %v", err)
			os.Exit(1)
		}

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = fmt.Sprintf(" AI agent is implementing the spec from %s...", filepath.Base(specPath))
		s.Start()

		// Delegate the entire implementation to gemini-cli, specifying the model.
		// The agent is responsible for all file I/O.
		agentOutput, err := ai.Executor("gemini", "-m", "gemini-2.5-pro", "-p", codePrompt)
		s.Stop()

		if err != nil {
			color.Red("Error during AI agent execution: %v", err)
			color.Yellow("Agent Output:\n%s", agentOutput)
			os.Exit(1)
		}

		// Print the final status update from the agent.
		color.Green("✔ Agent finished implementation.")
		fmt.Println(agentOutput)
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
}
