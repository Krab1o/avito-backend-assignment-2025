package model

type UserCreds struct {
	Username	string
	Password	string
}

type User struct {
	ID		int64
	Creds 	*UserCreds
	Coins 	int
}