package config

import "common/utils"

type GRPCConfig struct {
	Port string `mapstructure:"GRPC_PORT"`
}

func LoadGRPCConfig(path string) (*GRPCConfig, error) {
	viper, err := utils.GetViperInstance(path)
	if err != nil {
		return nil, err
	}
	cfg := &GRPCConfig{}
	viper.BindEnv("GRPC_PORT")
	viper.AutomaticEnv()
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
