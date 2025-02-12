package transaction

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

type transactionService struct {

}

func NewService() service.TransactionService {
	return &transactionService{}
}

//TODO: implement business-logic
//TODO: call repo-layer
func (s *transactionService) SendCoin(context.Context, *model.Transaction) error {
	return nil
}

//TODO: implement business-logic
//TODO: call repo-layer
func (s *transactionService) Info(context.Context) (*model.Info, error) {
	return &model.Info{}, nil
}