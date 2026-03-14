package main

import (
	"fmt"
	"net/http"
	"user-service/handler"
)

func main() {
	http.HandleFunc("/users", handler.CreateUserHandler)
	fmt.Println("User Service running on port 8081")

	http.ListenAndServe(":8081", nil)
}