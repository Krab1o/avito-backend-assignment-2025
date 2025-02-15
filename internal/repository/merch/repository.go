package merch

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/merch/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) repository.MerchRepository {
	return &repo{
		db: pool,
	}
}

func (r *repo) GetItem(ctx context.Context, tx pgx.Tx, itemTitle string) (*repoModel.Merch, error) {
	builder := sq.Select(
		repository.MerchIdColumn,
		repository.MerchTitleColumn,
		repository.MerchPriceColumn,
	).PlaceholderFormat(sq.Dollar).
	From(repository.MerchTableName).
	Where(sq.Eq{repository.MerchTitleColumn: itemTitle})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errs.NewDBError(err.Error(), err)
	}
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}
	merch := &repoModel.Merch{}
	err = row.Scan(
		&merch.ID,
		&merch.Title,
		&merch.Price,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, errs.NewDBError(err.Error(), err)
	}
	return merch, nil
}