package repository

import "wallet-service/model"

var wallets	= make(map[string]model.Wallet)
var transactions []model.Transaction

func SaveWallet(wallet model.Wallet){
	wallets[wallet.UserID] = wallet
}

func GetWallet(userID string)(model.Wallet, bool){
	wallet, exists := wallets[userID]
	return wallet, exists
}

func SaveTransaction(tx model.Transaction){
	transactions = append(transactions, tx)
}

func GetTransactions() []model.Transaction{
	return transactions
}
