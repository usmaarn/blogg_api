package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_DRIVER   string
	DB_SSL_MODE string
	DB_NAME     string
}

func loadEnv() (AppConfig, error) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, err
	}

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed", in.Name)
	})

	viper.WatchConfig()

	var cfg AppConfig

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return AppConfig{}, err
	}

	return cfg, nil
}
