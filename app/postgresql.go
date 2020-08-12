package app

import (
	"fmt"

	"github.com/caffeines/choto/config"
	"github.com/caffeines/choto/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var instance *gorm.DB

// ConnectSQLDB make SQL DB connection
func ConnectSQLDB() error {
	cnfg := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.DB().Username, config.DB().Password,
		config.DB().Host, config.DB().Port, config.DB().Name)

	db, err := gorm.Open("postgres", cnfg)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(config.DB().MaxIdleConnections)
	db.DB().SetMaxOpenConns(config.DB().MaxOpenConnections)
	db.DB().SetConnMaxLifetime(config.DB().MaxConnectionLifetime)

	db.LogMode(true)
	db.SetLogger(log.Log())

	instance = db

	return nil
}

// DB return database isntance
func DB() *gorm.DB {
	return instance
}
