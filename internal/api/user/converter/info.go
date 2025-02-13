package converter

import (
	dtoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/api/user/dto"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
)

func InfoDTOToService(dto *dtoModel.Info) *model.Info {
	// model.Inventory
	inventory := make([]model.Inventory, len(dto.Inventory))
	for i, item := range dto.Inventory {
		inventory[i] = model.Inventory{
			Type:     item.Type,
			Quantity: item.Quantity,
		}
	}
	// model.CoinHistory.Received
	received := make([]model.Received, len(dto.CoinHistory.Received))
	for i, r := range dto.CoinHistory.Received {
		received[i] = model.Received{
			FromUser: r.FromUser,
			Amount:   r.Amount,
		}
	}
	// model.CoinHistory.Sent
	sent := make([]model.Sent, len(dto.CoinHistory.Sent))
	for i, s := range dto.CoinHistory.Sent {
		sent[i] = model.Sent{
			ToUser: s.ToUser,
			Amount: s.Amount,
		}
	}
	return &model.Info{
		Coins:     dto.Coins,
		Inventory: inventory,
		CoinHistory: model.CoinHistory{
			Received: received,
			Sent:     sent,
		},
	}
}

func InfoServiceToDTO(model *model.Info) *dtoModel.Info {
	//dtoModel.Inventory
	inventory := make([]dtoModel.Inventory, len(model.Inventory))
	for i, item := range model.Inventory {
		inventory[i] = dtoModel.Inventory{
			Type:     item.Type,
			Quantity: item.Quantity,
		}
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