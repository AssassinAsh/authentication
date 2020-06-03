package server

import (
	"authentication/proto"
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
)

var client proto.AuthenticationServiceClient

func init() {
	conn, err := grpc.Dial("localhost:6565", grpc.WithInsecure())
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
		Username: "Rokade",
		Password: "abc123",
		Phone:    "9644695542",
	}

	res, _ := client.Register(context.Background(), &req)

	t.Log(res)
}
