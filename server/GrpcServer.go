package server

import (
	"authentication/client"
	"authentication/proto"
	"authentication/services"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userService services.UserServiceInterface
}

func (s *server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	return nil, nil
}

func (s *server) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	log.Println("Register rpc request :", request.String)
	res, err := s.userService.Register(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//StartServer - function to start the authentication grpc server
func StartServer() {
	listner, err := net.Listen("tcp", ":6565")

	if err != nil {
		log.Panic(err)
	}

	srv := grpc.NewServer()
	client.InitializeDB()

	proto.RegisterAuthenticationServiceServer(srv, &server{
		userService: services.NewUserService(),
	})

	reflection.Register(srv)

	if e := srv.Serve(listner); e != nil {
		log.Print(e)
	}
}
