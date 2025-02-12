package service

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

type TransactionService interface {
	SendCoin(context.Context, *model.Transaction) error
	Info(context.Context) (*model.Info, error)
}

type BuyingService interface {
	Buy(ctx context.Context, itemName string) error
}