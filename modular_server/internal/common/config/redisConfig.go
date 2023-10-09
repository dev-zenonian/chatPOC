package config

import "common/utils"

type RedisConfig struct {
	DSN string `mapstructure:"REDIS_DSN"`
}

func LoadRedisConfig(path string) (*RedisConfig, error) {
	viper, err := utils.GetViperInstance(path)
	if err != nil {
		return nil, err
	}
	cfg := &RedisConfig{}
	if err := viper.BindEnv("REDIS_DSN"); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil

}
