package converter

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/dto"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

func TransactionDTOToService(dto *dto.Transaction) *model.Transaction {
	return &model.Transaction{
		ToUser: dto.ToUser,
		Amount: dto.Amount,
	}
}

func TransactionServiceToDTO(model *model.Transaction) *dto.Transaction {
	return &dto.Transaction{
		ToUser: model.ToUser,
		Amount: model.Amount,
	}
}

