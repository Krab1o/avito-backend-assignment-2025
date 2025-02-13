package transaction

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) repository.TransactionRepository {
	return &repo{
		db: pool,
	}
}

func (r *repo) CreateTransaction(ctx context.Context, model *repoModel.Transaction) error {
	return nil
}
