package config

import "github.com/spf13/viper"

// appConfig 应用日志
type appConfig struct {
	AdminEmail string
}

func App() *appConfig {
	return &appConfig{
		AdminEmail: viper.GetString("app.adminEmail"),
	}
}
