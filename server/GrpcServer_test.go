package server

import (
	"authentication/config"
	"authentication/proto"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
)

var client proto.AuthenticationServiceClient

func init() {

	var cfg config.ServerConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	// target := cfg.Server.Host + ":" + cfg.Server.Port

	target := "localhost:6565"

	fmt.Println("Joining :", target)

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	client = proto.NewAuthenticationServiceClient(conn)

}

func TestLogin(t *testing.T) {

	req := proto.LoginRequest{
		Username: "Ashvin",
		Password: "abc123",
	}

	res, err := client.Login(context.Background(), &req)

	if res == nil {
		t.Errorf("Login failed %s", err.Error())
	} else {
		t.Logf("Login Successful : %s", res.String())
	}
}

func TestRegister(t *testing.T) {

	req := proto.RegisterRequest{
		Username: "vicky",
		Password: "ashvin786",
		Phone:    "9644695542",
	}

	res, _ := client.Register(context.Background(), &req)

	t.Log(res)
}
