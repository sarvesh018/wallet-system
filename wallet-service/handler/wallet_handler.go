package handler

import (
	"encoding/json"
	"net/http"
	"wallet-service/service"
)

type Request struct{
	UserID string 	`json:"user_id"`
	Amount int		`json:"amount"`
}

func CreateWalletHandler(w http.ResponseWriter, r *http.Request){
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	service.CreateWallet(req.UserID)
	
	w.Write([]byte("Wallet Created :)"))
}

func AddMoneyHandler(w http.ResponseWriter, r *http.Request){
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	service.AddMoney(req.UserID, req.Amount)
	w.Write([]byte("Money Added to the Wallet"))
}

func DeductMoneyHandler(w http.ResponseWriter, r *http.Request){
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	err := service.DeductMoney(req.UserID, req.Amount)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Money Deducted"))
}


