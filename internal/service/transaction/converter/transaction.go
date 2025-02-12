package converter

import (
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

func TransactionServiceToRepo(model *model.Transaction) *repoModel.Transaction {
	return &repoModel.Transaction{
		ToUser: model.ToUser,
		Amount: model.Amount,
	}
}

func TransactionRepoToService(repoModel *model.Transaction) *model.Transaction {
	return &model.Transaction{
		ToUser: repoModel.ToUser,
		Amount: repoModel.Amount,
	}
}