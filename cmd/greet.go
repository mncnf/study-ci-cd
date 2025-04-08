// Package cmd contains the implementation of the CLI commands for the application.
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{ //nolint:exhaustruct,gochecknoglobals
	Use: "greet",
	Run: func(_ *cobra.Command, args []string) {
		var name string

		if len(args) > 0 {
			name = args[0]
		}

		log.Println(greet(name))
	},
}

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(greetCmd)
}

func greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
