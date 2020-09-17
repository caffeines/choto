package migrations

import (
	"github.com/caffeines/choto/app"
	"github.com/caffeines/choto/core"
	"github.com/caffeines/choto/log"
	"github.com/caffeines/choto/models"
	"github.com/spf13/cobra"
)

//MigAutoCmd ...
var MigAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "auto alter database tables if required",
	Run:   auto,
}

func auto(cmd *cobra.Command, args []string) {
	tx := app.DB().Begin()

	var tables []core.Table
	tables = append(tables, &models.URL{})

	for _, t := range tables {
		if err := tx.AutoMigrate(t).Error; err != nil {
			tx.Rollback()
			log.Log().Errorln(err)
			return
		}
	}
	if err := tx.Commit().Error; err != nil {
		log.Log().Errorln(err)
		return
	}
	log.Log().Info("Migration auto completed")
}
