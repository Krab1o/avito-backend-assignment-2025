package transaction

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

type serv struct {
	transactionRepo repository.TransactionRepository
}

func NewService(tr repository.TransactionRepository) service.TransactionService {
	return &serv{
		transactionRepo: tr,
	}
}

//TODO: implement business-logic
//TODO: call repo-layer
func (s *serv) SendCoin(context.Context, *model.Transaction) error {
	return nil
}