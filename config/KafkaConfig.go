package config

//KafkaConfig - Kafka Configurations
type KafkaConfig struct {
	KafkaConfig struct {
		Producer struct {
			Server string `yaml:"bootstrap-server" env:"KAFKA_BOOTSTRAP_SERVER"`
			Topic  string `yaml:"topic"`
		} `yaml:"producer"`
		Consumer struct {
			Server string `yaml:"bootstrap-server" env:"KAFKA_BOOTSTRAP_SERVER"`
			Topic  string `yaml:"topic"`
		} `yaml:"consumer"`
	} `yaml:"kafka"`
}
