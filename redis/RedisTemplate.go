package redis

import (
	"authentication/connections"
	"fmt"
	"time"
)

//SaveToRedis - for saving in redis
func SaveToRedis(key string, value interface{}, eTime time.Duration) error {
	err := connections.RedisClient.Set(key, value, eTime).Err()

	if err != nil {
		panic(err)
	}

	fmt.Println("Saved to Redis")

	return nil
}
