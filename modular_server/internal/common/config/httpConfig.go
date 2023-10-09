package config

import "common/utils"

type HTTPConfig struct {
	Port string `mapstructure:"HTTP_PORT"`
}

func LoadHTTPConfig(path string) (*HTTPConfig, error) {
	viper, err := utils.GetViperInstance(path)
	if err != nil {
		return nil, err
	}
	cfg := &HTTPConfig{}
	viper.BindEnv("HTTP_PORT")
	viper.AutomaticEnv()
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
