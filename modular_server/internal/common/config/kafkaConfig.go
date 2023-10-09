package config

type KafkaConfig struct {
	Address   string `mapstructure:"KAFKA_ADDRESS"`
	Topic     string `mapstructure:"KAFKA_TOPIC"`
	Partition int    `mapstructure:"KAFKA_PARTITION"`
}

func LoadKafkaConfig(path string) (*KafkaConfig, error) {
	cfg := &KafkaConfig{
		Address:   "127.0.0.1:29092",
		Topic:     "notification",
		Partition: 0,
	}
	return cfg, nil
}
