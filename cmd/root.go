package cmd

import (
	"fmt"
	"os"

	"github.com/caffeines/choto/app"
	"github.com/caffeines/choto/config"
	"github.com/caffeines/choto/log"
	"github.com/spf13/cobra"
)

// RootCmd ...
var (
	RootCmd = &cobra.Command{
		Use:   "choto",
		Short: "A URL shortener http service",
		Long:  "A URL shortener REST API service",
	}
)

func init() {
	RootCmd.AddCommand(migrationCmd)
	RootCmd.AddCommand(serveCmd)
}

// Execute cmd...
func Execute() {
	log.SetupLog()

	if err := config.LoadConfig(); err != nil {
		fmt.Println("Failed to read config: ", err)
		os.Exit(1)
	}
	if err := app.ConnectSQLDB(); err != nil {
		fmt.Println("Failed to connect SQLDB: ", err)
		os.Exit(1)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
