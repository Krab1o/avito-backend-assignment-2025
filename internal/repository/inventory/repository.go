package inventory

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) repository.InventoryRepository {
	return &repo{
		db: pool,
	}
}

func (r *repo) BuyItem(ctx context.Context, buying *model.Buying) error {
	return nil
}