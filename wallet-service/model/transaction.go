package model

type Transaction struct{
	UserID string 	`json:"user_id"`
	Amount int		`json:"amount"`
	Type string 	`json:"type"`
}