package redis

import (
	"authentication/connections"
	"encoding/json"
	"fmt"
	"time"
)

//SaveToRedis - for saving in redis
func SaveToRedis(key string, value interface{}, eTime time.Duration) error {

	valueMarshalled, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	err = connections.RedisClient.Set(key, valueMarshalled, eTime).Err()

	if err != nil {
		panic(err)
	}

	fmt.Println("Saved to Redis")

	return nil
}

//UpdateRedisEntry - for updating entry in redis
func UpdateRedisEntry(key string, value interface{}) error {

	entry, err := connections.RedisClient.Get(key).Bytes()

	if err != nil {
		return err
	}

	ttl, err := connections.RedisClient.TTL(key).Result()

	if err != nil {
		return err
	}

	json.Unmarshal(entry, &value)

	SaveToRedis(key, value, ttl)

	return nil
}

//DeleteRedisEntry - deletes the redis entry at given key
func DeleteRedisEntry(key string) error {
	err := connections.RedisClient.Del(key)

	if err != nil {
		panic(err)
	}
	return err.Err()
}

//GetRedisEntry - to get entry saved in redis
func GetRedisEntry(key string) ([]byte, error) {
	return connections.RedisClient.Get(key).Bytes()
}
