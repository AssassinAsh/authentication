package main

import (
	"authentication/controller"
	"authentication/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Running...")
	http.HandleFunc("/status", controller.HealthCheckController)
	server.StartServer()
	fmt.Println(http.ListenAndServe(":8080", nil))
	fmt.Println("Started Authentication Server on Port : 8080")
}
