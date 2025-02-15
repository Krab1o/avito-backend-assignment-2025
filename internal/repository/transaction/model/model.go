package model

type Transaction struct {
	SenderID 	int64
	ReceiverID 	int64
	Amount 		int
}

type UserTransaction struct {
	Username string
	Amount int
}