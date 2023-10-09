package config

import "common/utils"

type GRPCClientConfig struct {
	Endpoint string `mapstructure:"GRPC_ENDPOINT"`
}

func LoadGRPCClientConfig(path string) (*GRPCClientConfig, error) {
	viper, err := utils.GetViperInstance(path)
	if err != nil {
		return nil, err
	}
	cfg := &GRPCClientConfig{}
	viper.BindEnv("GRPC_ENDPOINT")
	viper.AutomaticEnv()
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil

}
