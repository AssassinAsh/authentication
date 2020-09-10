package main

import (
	"authentication/connections"
	"authentication/models"
	"encoding/json"
	"fmt"
)

func main() {
	// // conn, err := grpc.Dial("localhost:6565", grpc.WithInsecure())
	// if err != nil {
	// 	log.Panic(err)
	// }
	// client := proto.NewAuthenticationServiceClient(conn)

	// test := otp{
	// 	Otp: *utils.OtpGenerator(),
	// }

	// fmt.Println("test :", test.Otp)

	// producers.OtpProducer(&test)

	// // conn.Close()

	value := models.OtpModel{}

	entry, err := connections.RedisClient.Get("test").Bytes()

	if err != nil {
		panic(err)
	}

	json.Unmarshal(entry, &value)

	fmt.Println(value.UserModel.Password)
}
