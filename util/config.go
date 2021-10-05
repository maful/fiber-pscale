package util

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) (err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	return nil
}
