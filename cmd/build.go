package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/fs"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A convenient wrapper for project-specific build commands.",
	Long:  "This command uses Go's os/exec package to call external commands defined in project-description.md or common build tools.",
	Run: func(cmd *cobra.Command, args []string) {
		buildCommand, err := fs.GetProjectCommand("build")
		if err != nil {
			color.Red("Error getting build command from project-description.md: %v", err)
			os.Exit(1)
		}

		color.Cyan("Executing build command: %s", buildCommand)
		buildCmd := exec.Command("bash", "-c", buildCommand)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		if err := buildCmd.Run(); err != nil {
			color.Red("Error executing build command: %v", err)
			os.Exit(1)
		}

		color.Green("Build command executed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}