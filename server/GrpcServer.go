package server

import (
	"authentication/config"
	"authentication/proto"
	"authentication/services"
	"context"
	"log"
	"net"

	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userService services.UserServiceInterface
}

func (s *server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	res, err := s.userService.Login(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *server) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	res, err := s.userService.Register(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *server) VerifyOtp(ctx context.Context, request *proto.OtpRequest) (*proto.OtpResponse, error) {
	res, err := s.userService.VerifyOtp(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//StartServer - function to start the authentication grpc server
func StartServer() {

	var cfg config.ServerConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		panic(err)
	}

	listner, err := net.Listen(cfg.Server.Network, ":"+cfg.Server.Port)

	if err != nil {
		log.Panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAuthenticationServiceServer(srv, &server{
		userService: services.NewUserService(),
	})

	reflection.Register(srv)

	if e := srv.Serve(listner); e != nil {
		log.Print(e)
	}
}
