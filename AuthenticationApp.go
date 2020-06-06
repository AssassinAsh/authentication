package main

import (
	"authentication/config"
	"authentication/controller"
	"authentication/kafka/consumers"
	"authentication/server"
	"fmt"
	"log"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jnewmano/grpc-json-proxy/gogoprotobuf/codec"
)

func main() {

	var cfg config.ServerConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		panic(err)
	}

	log.Println("Running...")
	http.HandleFunc("/status", controller.HealthCheckController)
	go consumers.OtpConsumer()
	server.StartServer()
	fmt.Println(http.ListenAndServe(":"+cfg.Server.Port, nil))
	fmt.Println("Started Authentication Server on Port :", cfg.Server.Port)
}
