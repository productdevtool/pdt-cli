package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/fs"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A convenient wrapper for project-specific deploy commands.",
	Long:  "This command uses Go's os/exec package to call external commands defined in project-description.md or common deploy tools.",
	Run: func(cmd *cobra.Command, args []string) {
		deployCommand, err := fs.GetProjectCommand("deploy")
		if err != nil {
			color.Red("Error getting deploy command from project-description.md: %v", err)
			os.Exit(1)
		}

		color.Cyan("Executing deploy command: %s", deployCommand)
		deployCmd := exec.Command("bash", "-c", deployCommand)
		deployCmd.Stdout = os.Stdout
		deployCmd.Stderr = os.Stderr

		if err := deployCmd.Run(); err != nil {
			color.Red("Error executing deploy command: %v", err)
			os.Exit(1)
		}

		color.Green("Deploy command executed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}