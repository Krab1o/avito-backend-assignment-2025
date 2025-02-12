package dto

type Transaction struct {
	ToUser string	`json:"toUser"`
	Amount int		`json:"amount"`
}