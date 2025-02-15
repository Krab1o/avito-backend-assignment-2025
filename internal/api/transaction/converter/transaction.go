package converter

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/dto"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

func TransactionDTOToService(dto *dto.Transaction, senderID int64) *model.Transaction {
	return &model.Transaction{
		FromUser: senderID,
		ToUser: dto.ToUser,
		Amount: dto.Amount,
	}
}

