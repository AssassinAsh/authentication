package models

//OtpKafka -
type OtpKafka struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Otp      string `json:"otp"`
}
