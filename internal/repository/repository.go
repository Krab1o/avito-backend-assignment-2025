package repository

import (
	"context"

	inventoryModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	transactionModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	userModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, model *transactionModel.Transaction) error
}

type InventoryRepository interface {
	BuyItem(ctx context.Context, buying *inventoryModel.Buying) error
}

type UserRepository interface {
	Info(ctx context.Context) (*userModel.Info, error)
	FindUserCreds(ctx context.Context, creds *userModel.UserCreds) (*userModel.UserCreds, error)
	CreateUser(ctx context.Context, creds *userModel.User) (int64, error)
}