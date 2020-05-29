package client

import (
	"log"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jinzhu/gorm"
)

//YAMLConfig - To parse application.yaml
type YAMLConfig struct {
	DBConfig struct {
		Username     string `yaml:"user"`
		Password     string `yaml:"password"`
		Port         string `yaml:"port"`
		DatabaseName string `yaml:"database_name"`
		Address      string `yaml:"address"`
	} `yaml:"database"`
}

//MysqlDB - function to initialize DB connection
func MysqlDB() {
	log.Print("Initializing DB Connection")

	var cfg YAMLConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", cfg.DBConfig.Username+":"+
		cfg.DBConfig.Password+"@tcp("+cfg.DBConfig.Address+":"+cfg.DBConfig.Port+")/"+
		cfg.DBConfig.DatabaseName)

	if err != nil {
		panic(err)
	}

	log.Print("DB Connection Successful")

	defer db.Close()
}
