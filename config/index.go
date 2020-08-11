package config

import (
	"sync"

	"github.com/spf13/viper"
)

var mu sync.Mutex

// LoadConfig load all configuration
func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	LoadApp()
	return nil
}
