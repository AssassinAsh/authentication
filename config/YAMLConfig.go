package config

//YAMLConfig - To parse application.yaml
type YAMLConfig struct {
	DBConfig struct {
		Username     string `yaml:"user"`
		Password     string `yaml:"password"`
		Port         string `yaml:"port"`
		DatabaseName string `yaml:"database_name"`
		Address      string `yaml:"address"`
	} `yaml:"database"`
	RedisConfig struct {
		Env     string `yaml:"env"`
		Port    string `yaml:"port"`
		Address string `yaml:"address"`
	} `yaml:"redis"`
}
