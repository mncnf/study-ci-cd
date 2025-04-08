package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //nolint:exhaustruct,gochecknoglobals
	Use: "study-ci-cd",
	Run: func(_ *cobra.Command, _ []string) {},
}

// Execute runs the root command for the application.
// If an error occurs, the program exits with a status code of 1.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
