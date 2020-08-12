package migrations

import (
	"fmt"

	"github.com/spf13/cobra"
)

var MigAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "auto alter database tables if required",
	Run:   auto,
}

func auto(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
