package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{ //nolint:exhaustruct,gochecknoglobals
	Use: "serve",
	Run: func(_ *cobra.Command, _ []string) {
		e := echo.New()

		e.Use(middleware.Logger())

		e.Static("/", "web")

		e.Logger.Fatal(e.Start(":8080"))
	},
}

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(serveCmd)
}
