package constants

import "time"

//OtpLength - Length of otp
var OtpLength = 4

//RedisExpTime - 15 min time for redis entry
var RedisExpTime time.Duration = 15

//OtpMessage - message for otp kafka
var OtpMessage string = "The otp raised by kafka."
