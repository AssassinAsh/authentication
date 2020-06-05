package connections

import (
	"authentication/config"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jinzhu/gorm"
)

//DBInstanceInterface -
type DBInstanceInterface interface {
	GetDatabaseInstance() *gorm.DB
}

//DBInstance -
type DBInstance struct {
	DBInstanceInterface
}

type connectionDB struct {
	dataBase *gorm.DB
}

var db *connectionDB
var onceDB sync.Once

func getDBConfig() *connectionDB {
	onceDB.Do(func() {
		db = &connectionDB{
			dataBase: getMysqlDB(),
		}
	})
	return db
}

//InitializeDB - Initializing DB
func InitializeDB() {
	getDBConfig()
}

//MysqlDB - function to initialize DB connection
func getMysqlDB() *gorm.DB {
	log.Print("Initializing DB Connection")

	var cfg config.DBConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", cfg.DBConfig.Username+":"+
		cfg.DBConfig.Password+"@tcp("+cfg.DBConfig.Address+":"+cfg.DBConfig.Port+")/"+
		cfg.DBConfig.DatabaseName+"?parseTime=true")

	if err != nil {
		panic(err)
	}

	log.Print("DB Connection Successful")

	return db
}

//GetDatabaseInstance -
func (db *DBInstance) GetDatabaseInstance() *gorm.DB {
	return getDBConfig().dataBase
}
