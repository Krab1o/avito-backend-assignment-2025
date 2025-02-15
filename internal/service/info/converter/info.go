package converter

import (
	repoInv "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	repoTr "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
)

func InfoRepoToService(
	coins int,
	sendings []repoTr.UserTransaction,
	receivings []repoTr.UserTransaction,
	inventory repoInv.Inventory,
) *model.Info {
	serviceReceived := make([]model.Received, len(receivings))
	for i, val := range receivings {
		serviceReceived[i] = model.Received{
			FromUser: val.Username,
			Amount: val.Amount,
		}
	}
	serviceSent := make([]model.Sent, len(sendings))
	for i, val := range sendings {
		serviceSent[i] = model.Sent{
			ToUser: val.Username,
			Amount: val.Amount,
		}
	}
	return &model.Info{
		Coins: coins,
		CoinHistory: model.CoinHistory{
			Received: serviceReceived,
			Sent: serviceSent,
		},
		Inventory: model.Inventory(inventory),
	}
}

// func InfoServiceToRepo(model *model.Info) *repoModel.Info {
// 	inventory := make([]repoModel.Inventory, len(model.Inventory))
// 	for i, item := range model.Inventory {
// 		inventory[i] = repoModel.Inventory{
// 			Type:     item.Type,
// 			Quantity: item.Quantity,
// 		}
// 	}

// 	received := make([]repoModel.Received, len(model.CoinHistory.Received))
// 	for i, r := range model.CoinHistory.Received {
// 		received[i] = repoModel.Received{
// 			FromUser: r.FromUser,
// 			Amount:   r.Amount,
// 		}
// 	}

// 	sent := make([]repoModel.Sent, len(model.CoinHistory.Sent))
// 	for i, s := range model.CoinHistory.Sent {
// 		sent[i] = repoModel.Sent{
// 			ToUser: s.ToUser,
// 			Amount: s.Amount,
// 		}
// 	}

// 	return &repoModel.Info{
// 		Coins:     model.Coins,
// 		Inventory: inventory,
// 		CoinHistory: repoModel.CoinHistory{
// 			Received: received,
// 			Sent:     sent,
// 		},
// 	}
// }

// func InfoRepoToService(repoModel *repoModel.Info) *model.Info {
// 	inventory := make([]model.Inventory, len(repoModel.Inventory))
// 	for i, item := range repoModel.Inventory {
// 		inventory[i] = model.Inventory{
// 			Type:     item.Type,
// 			Quantity: item.Quantity,
// 		}
// 	}

// 	received := make([]model.Received, len(repoModel.CoinHistory.Received))
// 	for i, r := range repoModel.CoinHistory.Received {
// 		received[i] = model.Received{
// 			FromUser: r.FromUser,
// 			Amount:   r.Amount,
// 		}
// 	}

// 	sent := make([]model.Sent, len(repoModel.CoinHistory.Sent))
// 	for i, s := range repoModel.CoinHistory.Sent {
// 		sent[i] = model.Sent{
// 			ToUser: s.ToUser,
// 			Amount: s.Amount,
// 		}
// 	}

// 	return &model.Info{
// 		Coins:     repoModel.Coins,
// 		Inventory: inventory,
// 		CoinHistory: model.CoinHistory{
// 			Received: received,
// 			Sent:     sent,
// 		},
// 	}
// }