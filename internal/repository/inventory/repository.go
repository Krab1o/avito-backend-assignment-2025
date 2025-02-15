package inventory

import (
	"context"
	"fmt"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
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

func (r *repo) BuyItem(ctx context.Context, tx pgx.Tx, buying *model.Buying) error {
	builder := sq.Insert(repository.InventoryTableName).
        PlaceholderFormat(sq.Dollar).
        Columns(
			repository.InventoryIdMerchColumn,
			repository.InventoryIdUserColumn,
			repository.InventoryQuantityColumn,
		).Values(
			buying.MerchID,
            buying.BuyerID,
            buying.Quantity,
		).Suffix(fmt.Sprintf(
            "ON CONFLICT (%s, %s) DO UPDATE SET %s = %s.%s + %d",
            repository.InventoryIdUserColumn,
            repository.InventoryIdMerchColumn,
            repository.InventoryQuantityColumn,
            repository.InventoryTableName,
            repository.InventoryQuantityColumn,
            buying.Quantity,
        ))
    query, row, err := builder.ToSql()
    if err != nil {
        return errs.NewDBError(err.Error(), err)
    }
    if tx != nil {
        _, err = tx.Exec(ctx, query, row...)
    } else {
        _, err = r.db.Exec(ctx, query, row...)
    }
    if err != nil {
        return errs.NewDBError(err.Error(), err)
    }
    return nil
}
