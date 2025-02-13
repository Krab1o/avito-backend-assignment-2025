package model

type User struct {
	ID				int64
	Creds 			UserCreds
	Coins 			int
}

type UserCreds struct {
	Username		string
	PasswordHash	string
}