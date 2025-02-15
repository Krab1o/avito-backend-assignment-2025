package service

import (
	"context"

	userModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
	buyingModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
	infoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
	transactionModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
)

type AuthService interface {
	Auth(context.Context, *userModel.UserCreds) (string, error)
}

type InfoService interface {
	Info(context.Context, int64) (*infoModel.Info, error)
}

type TransactionService interface {
	SendCoin(context.Context, *transactionModel.Transaction) error
}

type BuyingService interface {
	Buy(context.Context, *buyingModel.Buying) error
}