package services

import (
	"authentication/proto"
	"authentication/repo/master"
	"fmt"
	"log"

	"authentication/mappers"
)

//UserServiceInterface -
type UserServiceInterface interface {
	Login(request *proto.LoginRequest) (*proto.LoginResponse, error)
	Register(request *proto.RegisterRequest) (*proto.RegisterResponse, error)
}

//UserService -
type UserService struct {
	userCredentialsMasterRepoService master.UserCredentialsMasterRepoService
}

//NewUserService -
func NewUserService() UserServiceInterface {
	return &UserService{
		userCredentialsMasterRepoService: master.NewUserCredentialsMasterRepo(),
	}
}

//Login -
func (user *UserService) Login(request *proto.LoginRequest) (*proto.LoginResponse, error) {
	return nil, nil
}

//Register -
func (user *UserService) Register(request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	log.Println("userService")
	masterRepo := user.userCredentialsMasterRepoService
	err := masterRepo.SaveToUserCredentials(mappers.RegisterRequestProtoToUserModel(request))
	if err != nil {
		fmt.Println(err)
	}

	res := proto.RegisterResponse{RegisterResponseMap: map[string]string{
		"Response Code" : "200"
	}}

	return &res, err
}
