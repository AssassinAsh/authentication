package server

import (
	"authentication/config"
	"authentication/proto"
	"authentication/services"
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
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

	fmt.Println("Starting gRPC Server")

	var cfg config.ServerConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		panic(err)
	}

// 	listner, err := net.Listen(cfg.Server.Network, ":"+cfg.Server.Port)

	if err != nil {
		log.Panic(err)
	}

	srv := grpc.NewServer()

    proto.RegisterAuthenticationServiceServer(srv, &server{
		userService: services.NewUserService(),
	})

//     wrappedServer := grpcweb.WrapServer(srv)

    wrappedGrpc := grpcweb.WrapServer(srv, grpcweb.WithAllowedRequestHeaders([]string{"*"}), grpcweb.WithWebsockets(false))


// 	handler := func(resp http.ResponseWriter, req *http.Request) {
// 		wrappedServer.ServeHTTP(resp, req)
// 	}

    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    		w.Header().Set("Access-Control-Allow-Origin", "*")
    		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-user-agent, x-grpc-web, grpc-status, grpc-message")
    		w.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
    		if r.Method == "OPTIONS" {
    			return
    		}

    		if wrappedGrpc.IsGrpcWebRequest(r) {
    			wrappedGrpc.ServeHTTP(w, r)
    		} else {
    			// Fall back to other servers.
    			http.DefaultServeMux.ServeHTTP(w, r)
    		}
    	})

	port, _ := strconv.Atoi(cfg.Server.Port)


	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}


	fmt.Printf("Starting server. http port: %d", port)

	if err := httpServer.ListenAndServe(); err != nil {
			fmt.Sprintf("failed starting http server: %v", err)
	}
//
// 	reflection.Register(srv)
//
// 	if e := srv.Serve(listner); e != nil {
// 		log.Print(e)
// 	}
}
