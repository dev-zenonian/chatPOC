package config

import "common/utils"

type MongoConfig struct {
	DSN      string `mapstructure:"MONGO_DSN"`
	Database string `mapstructure:"MONGO_DATABASE"`
}

func LoadMongoConfig(path string) (*MongoConfig, error) {
	viper, err := utils.GetViperInstance(path)
	if err != nil {
		return nil, err
	}
	cfg := &MongoConfig{}
	if err := viper.BindEnv("MONGO_DSN"); err != nil {
		return nil, err
	}

	if err := viper.BindEnv("MONGO_DATABSE"); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
