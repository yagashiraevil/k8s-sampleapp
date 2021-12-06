package util

import "github.com/spf13/viper"

type Config struct {
	BindAddr string `mapstructure:"bind_addr"`
	BindPort string `mapstructure:"bind_port"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return
}
