package config

//ServerConfig - Go Server Configurations
type ServerConfig struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Network string `yaml:"network"`
	} `yaml:"server"`
}
