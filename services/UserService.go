package services

import (
	"authentication/proto"
	"authentication/repo/master"
	"authentication/repo/slave"
	"fmt"
	"log"

	"authentication/mappers"
	"authentication/utils"
)

//UserServiceInterface -
type UserServiceInterface interface {
	Login(request *proto.LoginRequest) (*proto.LoginResponse, error)
	Register(request *proto.RegisterRequest) (*proto.RegisterResponse, error)
}

//UserService -
type UserService struct {
	userCredentialsMasterRepoService master.UserCredentialsMasterRepoService
	userCredentialsSlaveRepoService  slave.UserCredentialsSlaveRepoService
}

//NewUserService -
func NewUserService() UserServiceInterface {
	return &UserService{
		userCredentialsMasterRepoService: master.NewUserCredentialsMasterRepo(),
		userCredentialsSlaveRepoService:  slave.NewUserCredentialsSlaveRepo(),
	}
}

//Login -
func (user *UserService) Login(request *proto.LoginRequest) (*proto.LoginResponse, error) {
	log.Println("User Login Service")
	slaveRepo := user.userCredentialsSlaveRepoService
	userModel, err := slaveRepo.ReadUserCredentialsByUsername(request.Username)

	if err != nil {
		return &proto.LoginResponse{
			LoginResponseMap: map[string]string{
				"Response":      "Error in authentication",
				"Response Code": "404",
			},
		}, err
	}

	if utils.ComparePasswords(userModel.Password, request.Password+request.Username) {
		return &proto.LoginResponse{
			LoginResponseMap: map[string]string{
				"Response":      "Login Successful",
				"Response Code": "200",
			},
		}, nil
	}
	return &proto.LoginResponse{
		LoginResponseMap: map[string]string{
			"Response":      "Wrong Password",
			"Response Code": "200",
		},
	}, nil

}

//Register -
func (user *UserService) Register(request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	log.Println("User Register Service")
	masterRepo := user.userCredentialsMasterRepoService

	if user.checkUsernameExistence(request.Username) {

		request.Password = utils.HashAndSaltPassword(request.Password + request.Username)

		err := masterRepo.SaveToUserCredentials(mappers.RegisterRequestProtoToUserModel(request))
		if err != nil {
			fmt.Println(err)
		}

		res := proto.RegisterResponse{RegisterResponseMap: map[string]string{
			"Response":      "Account Created Successfully",
			"Response Code": "200",
		}}

		return &res, err
	}

	res := proto.RegisterResponse{RegisterResponseMap: map[string]string{
		"Response":      "Account Already Exists",
		"Response Code": "200",
	}}

	return &res, nil

}

func (user *UserService) checkUsernameExistence(username string) bool {
	slaveRepo := user.userCredentialsSlaveRepoService

	_, err := slaveRepo.ReadUserCredentialsByUsername(username)

	if err != nil {
		log.Println(err)
		return true
	}

	return false
}
