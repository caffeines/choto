package cmd

import (
	"github.com/caffeines/choto/cmd/migrations"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration migrates database schemas",
}

func init() {
	migrationCmd.AddCommand(migrations.MigAutoCmd)
}
