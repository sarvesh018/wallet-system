package main

import(
	"fmt"
	"net/http"
	"wallet-service/handler"
)

func main() {
	http.HandleFunc("/wallet/create", handler.CreateWalletHandler)
	http.HandleFunc("/wallet/add", handler.AddMoneyHandler)
	http.HandleFunc("/wallet/deduct", handler.DeductMoneyHandler)

	fmt.Println("Wallet service is running on port 8082")
	http.ListenAndServe(":8082", nil)
}