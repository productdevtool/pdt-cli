package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdt",
	Short: "pdt is the command center for the AI-native founder",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application. For example: ...`,
	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
