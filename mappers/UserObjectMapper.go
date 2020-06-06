package mappers

import (
	"authentication/models"
	"authentication/proto"
)

//RegisterRequestProtoToUserModel - mapper
func RegisterRequestProtoToUserModel(request *proto.RegisterRequest) *models.User {
	return &models.User{
		Username: request.Username,
		Password: request.Password,
		Phone:    request.Phone,
	}
}

//UserToOtpModel -
func UserToOtpModel(user *models.User) *models.OtpModel {
	return &models.OtpModel{
		UserModel: *user,
		Otp:       "",
	}
}
