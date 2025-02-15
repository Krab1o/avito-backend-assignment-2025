package model

type Info struct {
	Coins       int         
	Inventory   Inventory
	CoinHistory CoinHistory
}
type Inventory map[string]int

type Received struct {
	FromUser string
	Amount   int
}
type Sent struct {
	ToUser string
	Amount int
}
type CoinHistory struct {
	Received []Received 
	Sent     []Sent 
}