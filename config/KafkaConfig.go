package config

//KafkaConfig - Kafka Configurations
type KafkaConfig struct {
	KafkaConfig struct {
		Producer struct {
			Otp struct {
				Server string `yaml:"bootstrap-server" env:"KAFKA_BOOTSTRAP_SERVER"`
				Topic  string `yaml:"topic"`
			} `yaml:"otp"`
		} `yaml:"producer"`
		Consumer struct {
			Otp struct {
				Server string `yaml:"bootstrap-server" env:"KAFKA_BOOTSTRAP_SERVER"`
				Topic  string `yaml:"topic"`
				Group  string `yaml:"group"`
			} `yaml:"otp"`
		} `yaml:"consumer"`
	} `yaml:"kafka"`
}
