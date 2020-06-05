package main

import (
	"authentication/kafka/producers"
	"authentication/utils"
	"fmt"
)

type otp struct {
	Otp string
}

func main() {
	// // conn, err := grpc.Dial("localhost:6565", grpc.WithInsecure())
	// if err != nil {
	// 	log.Panic(err)
	// }
	// client := proto.NewAuthenticationServiceClient(conn)

	test := otp{
		Otp: *utils.OtpGenerator(),
	}

	fmt.Println("test :", test.Otp)

	producers.OtpProducer(&test)

	// conn.Close()

}
