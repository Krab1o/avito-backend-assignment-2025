package transaction

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5"
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

func (r *repo) CreateTransaction(ctx context.Context, tx pgx.Tx, model *repoModel.Transaction) error {
	builder := sq.Insert(repository.TransactionTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			repository.TransactionIdSenderColumn,
			repository.TransactionIdReceiverColumn,
			repository.TransactionAmountColumn,
		).Values(
			model.SenderID, 
			model.ReceiverID, 
			model.Amount,
		)
	query, args, err := builder.ToSql()
	if err != nil {
		errs.NewDBError(err.Error(), err)
	}
	if tx != nil {
		_, err = tx.Exec(ctx, query, args...)
	} else {
		_, err = r.db.Exec(ctx, query, args...)	
	}
	if err != nil {
		return errs.NewDBError(err.Error(), err)
	}
	return nil
}
