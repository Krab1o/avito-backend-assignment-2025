package repository

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	inventoryModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	merchModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/merch/model"
	transactionModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	userModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/jackc/pgx/v5"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, tx pgx.Tx, model *transactionModel.Transaction) error
	GetTransactionsByReceiverID(ctx context.Context, tx pgx.Tx, receiverID int64) ([]transactionModel.UserTransaction, error)
	GetTransactionsBySenderID(ctx context.Context, tx pgx.Tx, senderID int64) ([]transactionModel.UserTransaction, error)
}

type InventoryRepository interface {
	GetInventoryByID(ctx context.Context, tx pgx.Tx, ownerID int64) (model.Inventory, error)
	BuyItem(ctx context.Context, tx pgx.Tx, buying *inventoryModel.Buying) error
}

type UserRepository interface {
	WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error
	GetUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*userModel.User, error)
	GetUserByID(ctx context.Context, tx pgx.Tx, id int64) (*userModel.User, error)
	CreateUser(ctx context.Context, tx pgx.Tx, creds *userModel.User) (int64, error)
	AddCoins(ctx context.Context, tx pgx.Tx, user *userModel.User, value int) error
	SubtractCoins(ctx context.Context, tx pgx.Tx, user *userModel.User, value int) error
}

type MerchRepository interface {
	GetItem(ctx context.Context, tx pgx.Tx, itemTitle string) (*merchModel.Merch, error)
}