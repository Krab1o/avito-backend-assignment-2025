package converter

import (
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

func TransactionServiceToRepo(model *model.Transaction, receiverID int64) *repoModel.Transaction {
	return &repoModel.Transaction{
		SenderID: model.FromUser,
		ReceiverID: receiverID,
		Amount: model.Amount,
	}
}