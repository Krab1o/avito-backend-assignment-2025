package converter

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
)

func NewBuying(merchID int64, userID int64, quantity int) *model.Buying {
	return &model.Buying{
		BuyerID: userID,
		MerchID: merchID,
		Quantity: quantity,
	}
}