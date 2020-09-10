package models

//OtpModel -
type OtpModel struct {
	UserModel User   `json:"user"`
	Otp       string `json:"otp"`
}
