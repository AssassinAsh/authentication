package components

import (
	"authentication/constants"
	"authentication/kafka/producers"
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
func OtpSender(username string) {

	otp := models.OtpKafka{
		Message:  constants.OtpMessage,
		Username: username,
		Otp:      *utils.OtpGenerator(),
	}

	producers.OtpProducer(&otp)
}

//CacheUser -
func CacheUser(user *models.User) error {
	expTime := getRedisExpTime(&constants.RedisExpTime)

	otpModel := mappers.UserToOtpModel(user)

	err := redis.SaveToRedis(user.Username, &otpModel, expTime.Sub(time.Now()))

	return err
}

func getRedisExpTime(time *time.Duration) time.Time {
	expTime := utils.GetTimeAfterDuration(time)

	return utils.ConvertToUTC(expTime)
}

//UpdateOtp - for updating the consumed otp into redis
func UpdateOtp(otp *models.OtpKafka) error {

	otpModel := models.OtpModel{}

	entry, err := redis.GetRedisEntry(otp.Username)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(entry, &otpModel)

	otpModel.Otp = otp.Otp

	return redis.UpdateRedisEntry(otp.Username, &otpModel)
}
