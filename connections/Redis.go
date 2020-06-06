package connections

import (
	"authentication/config"
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/ilyakaznacheev/cleanenv"
)

//RedisClient - redis client
var RedisClient *redis.Client

//Redis - function to connect to redis
func Redis() {
	fmt.Println("Initializing Redis...")

	var cfg config.RedisConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		panic(err)
	}

	dsn := os.Getenv(cfg.RedisConfig.Env)

	if len(dsn) == 0 {
		dsn = cfg.RedisConfig.Address + ":" + cfg.RedisConfig.Port
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	_, err = RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}
