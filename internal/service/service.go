package service

import (
	"context"

	buyingModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
	transactionModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

type TransactionService interface {
	SendCoin(context.Context, *transactionModel.Transaction) error
	Info(context.Context) (*transactionModel.Info, error)
}

type BuyingService interface {
	Buy(context.Context, *buyingModel.Buying) error
}