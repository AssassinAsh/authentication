package redis

import (
	"authentication/connections"
	"time"
)

//SaveToRedis - for saving in redis
func SaveToRedis(key string, value interface{}, eTime *time.Duration) error {
	err := connections.RedisClient.Set(key, value, *eTime).Err()

	if err != nil {
		panic(err)
	}

	return nil
}
