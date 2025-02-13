package service

import (
	"context"

	buyingModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
	infoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
	transactionModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

type AuthService interface {
	Auth(context.Context) error
}

type InfoService interface {
	Info(context.Context) (*infoModel.Info, error)
}

type TransactionService interface {
	SendCoin(context.Context, *transactionModel.Transaction) error
}

type BuyingService interface {
	Buy(context.Context, *buyingModel.Buying) error
}