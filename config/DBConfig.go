package config

//DBConfig - Database Configurations
type DBConfig struct {
	DBConfig struct {
		Username     string `yaml:"user"`
		Password     string `yaml:"password"`
		Port         string `yaml:"port"`
		DatabaseName string `yaml:"database_name"`
		Address      string `yaml:"address"`
	} `yaml:"database"`
}
