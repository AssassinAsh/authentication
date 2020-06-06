package components

import (
	"authentication/constants"
	"authentication/mappers"
	"authentication/models"
	"authentication/redis"
	"authentication/repo/slave"
	"authentication/utils"
	"encoding/json"
	"log"
	"time"
)

//CheckUsernameExistence -
func CheckUsernameExistence(username string) bool {
	slaveRepo := slave.NewUserCredentialsSlaveRepo()

	_, err := slaveRepo.ReadUserCredentialsByUsername(username)

	if err != nil {
		log.Println(err)
		return true
	}

	return false
}

//OtpSender -
func OtpSender() error {

	return nil
}

//CacheUser -
func CacheUser(user *models.User) error {
	expTime := getRedisExpTime(&constants.RedisExpTime)

	otpModel, err := json.Marshal(mappers.UserToOtpModel(user))

	if err != nil {
		panic(err)
	}

	err = redis.SaveToRedis(user.Username, otpModel, expTime.Sub(time.Now()))

	return err
}

func getRedisExpTime(time *time.Duration) time.Time {
	expTime := utils.GetTimeAfterDuration(time)

	return utils.ConvertToUTC(expTime)
}
