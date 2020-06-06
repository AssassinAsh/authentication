package utils

import "time"

//GetStringHash -
func GetStringHash(input string) []byte {
	return []byte(input)
}

//GetTimeAfterDuration - to get time after input duration
func GetTimeAfterDuration(input time.Duration) int64 {
	return time.Now().Add(time.Minute * input).Unix()
}

//ConvertToUTC - for converting Unix to UTC(to Time object)
func ConvertToUTC(input int64) time.Time {
	return time.Unix(input, 0)
}
