package main

import (
	"authentication/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6565", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	client := proto.NewAuthenticationServiceClient(conn)

	req := proto.RegisterRequest{
		Username: "Rokade",
		Password: "abc124",
		Phone:    "9644695542",
	}

	res, err := client.Register(context.Background(), &req)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(res)
	}

	conn.Close()
}
