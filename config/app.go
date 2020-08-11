package config

import "github.com/spf13/viper"

// Application holds the application configuration
type Application struct {
	Base      string
	Port      int
	JWTSecret string
}

var app Application

// App returns the default application configuration
func App() *Application {
	return &app
}

// LoadApp loads application configuration
func LoadApp() {
	app = Application{
		Base:      viper.GetString("app.host"),
		Port:      viper.GetInt("app.port"),
		JWTSecret: viper.GetString("app.jwtKey"),
	}
}
