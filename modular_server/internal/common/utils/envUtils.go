package utils

import (
	"github.com/spf13/viper"
)

func GetViperInstance(path string) (*viper.Viper, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return viper.GetViper(), nil
}
