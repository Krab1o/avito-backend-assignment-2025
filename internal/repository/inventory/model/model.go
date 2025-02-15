package model

type Buying struct {
	ID		 int64
	BuyerID  int64
	MerchID  int64
	Quantity int
}

type Inventory map[string]int