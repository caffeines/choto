package migrations

import (
	"fmt"

	"github.com/caffeines/choto/log"
	"github.com/spf13/cobra"
)

//MigAutoCmd ...
var MigAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "auto alter database tables if required",
	Run:   auto,
}

func auto(cmd *cobra.Command, args []string) {
	fmt.Println(args)
	log.Log().Info("Migration auto completed")
}
