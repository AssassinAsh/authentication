package config

//RedisConfig - Redis Configurations
type RedisConfig struct {
	RedisConfig struct {
		Env     string `yaml:"env"`
		Port    string `yaml:"port"`
		Address string `yaml:"address"`
	} `yaml:"redis"`
}
