package clients

// import (
// 	"fmt"
// 	"os"

// 	"github.com/go-redis/redis/v7"
// )

// var client *redis.Client

// //Init - function to connect to redis
// func Redis() {
// 	fmt.Println("Initializing Redis...")
// 	dsn := os.Getenv("REDIS_DSN")
// 	if len(dsn) == 0 {
// 		dsn = "localhost:6379"
// 	}
// 	client = redis.NewClient(&redis.Options{
// 		Addr: dsn,
// 	})
// 	_, err := client.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	}
// }
