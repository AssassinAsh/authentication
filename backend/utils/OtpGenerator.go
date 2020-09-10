package utils

import (
	"authentication/constants"
	"math/rand"
	"strconv"
	"time"
)

//OtpGenerator - To generate a random otp
func OtpGenerator() *string {
	var otp string
	for i := 0; i < constants.OtpLength; i++ {
		otp += strconv.Itoa(randomInt(0, 9))
	}
	return &otp
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
