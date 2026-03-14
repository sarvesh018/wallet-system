package service

import (
	"errors"
	"wallet-service/model"
	"wallet-service/repository"
)

func CreateWallet(userID string){
	wallet := model.Wallet{
		UserID: userID,
		Balance: 0,
	}
	repository.SaveWallet(wallet)
}

func AddMoney(userID string, amount int){
	wallet, _ := repository.GetWallet(userID)
	wallet.Balance += amount

	repository.SaveWallet(wallet)

	tx := model.Transaction{
		UserID: userID,
		Amount: amount,
		Type: "credit",
	}
	repository.SaveTransaction(tx)
}

func DeductMoney(userID string, amount int) error {
	wallet, exists := repository.GetWallet(userID)

	if !exists{
		return errors.New("Wallet not found")
	}
	if wallet.Balance < amount {
		return errors.New("Insufficient Balance :(")
	}

	wallet.Balance -= amount
	repository.SaveWallet(wallet)

	tx := model.Transaction{
		UserID: userID,
		Amount: amount,
		Type: "debit",
	}
	repository.SaveTransaction(tx)
	return nil
}