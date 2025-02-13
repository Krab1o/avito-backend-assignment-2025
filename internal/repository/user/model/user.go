package model

type User struct {
	ID				int64
	Username 		string
	PasswordHash 	string
	Coins 			int
}

type UserCreds struct {
	Username		string
	PasswordHash	string
}