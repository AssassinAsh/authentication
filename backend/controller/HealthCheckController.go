package controller

import (
	"fmt"
	"net/http"
)

//HealthCheckController - to checkt the status of the server
func HealthCheckController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Status OK...")
	w.Write([]byte("Status Checked"))
}
