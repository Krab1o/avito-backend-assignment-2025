package converter

import (
	dtoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/api/user/dto"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
)

func InfoServiceToDTO(model *model.Info) *dtoModel.Info {
	//dtoModel.Inventory
	inventory := make([]dtoModel.Inventory, len(model.Inventory))
	i := 0
	for title, quantity := range model.Inventory {
		inventory[i] = dtoModel.Inventory{
			Type:     title,
			Quantity: quantity,
		}
		i++
	}
	// dtoModel.CoinHistory.Received
	received := make([]dtoModel.Received, len(model.CoinHistory.Received))
	for i, r := range model.CoinHistory.Received {
		received[i] = dtoModel.Received{
			FromUser: r.FromUser,
			Amount:   r.Amount,
		}
	}
	// dtoModel.CoinHistory.Sent
	sent := make([]dtoModel.Sent, len(model.CoinHistory.Sent))
	for i, s := range model.CoinHistory.Sent {
		sent[i] = dtoModel.Sent{
			ToUser: s.ToUser,
			Amount: s.Amount,
		}
	}
	return &dtoModel.Info{
		Coins:		model.Coins,
		Inventory: 	inventory,
		CoinHistory: dtoModel.CoinHistory{
			Received: received,
			Sent:     sent,
		},
	}
}