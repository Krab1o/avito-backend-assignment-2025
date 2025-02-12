package transaction

import (
	"context"

	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
)

type transactionRepository interface {
	Create(ctx context.Context, model *repoModel.Transaction)
	Info(ctx context.Context, model *repoModel.Info)
}