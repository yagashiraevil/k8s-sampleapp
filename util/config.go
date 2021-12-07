package util

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	BindAddr string `mapstructure:"bind_addr"`
	DBHost   string `mapstructure:"db_host"`
	DBPort   string `mapstructure:"db_port"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			config.BindAddr = os.Getenv("bind_addr")
			config.DBHost = os.Getenv("db_host")
			config.DBPort = os.Getenv("db_port")
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return
}
