package transaction

import (
	"context"
	"fmt"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5"
)

func (r *repo) getWithRequest(ctx context.Context, tx pgx.Tx, builder sq.SelectBuilder) ([]repoModel.UserTransaction, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errs.NewDBError(err.Error(), err)
	}
	var rows pgx.Rows
	if tx != nil {
		rows, err = tx.Query(ctx, query, args...)
	} else {
		rows, err = r.db.Query(ctx, query, args...)
	}
	if err != nil {
		return nil, errs.NewDBError(err.Error(), err)
	}
	var transactionArr []repoModel.UserTransaction
	transaction := repoModel.UserTransaction{}
	for rows.Next() {
		err = rows.Scan(
			&transaction.Username,
			&transaction.Amount,
		)
		transactionArr = append(transactionArr, transaction)
	}
	if err != nil {
		return nil, errs.NewDBError(err.Error(), err)
	}
	return transactionArr, nil
}

func (r *repo) GetTransactionsByReceiverID(ctx context.Context, tx pgx.Tx, receiverID int64) ([]repoModel.UserTransaction, error) {
	builder := sq.Select(
			repository.UserUsernameColumn,
			repository.TransactionAmountColumn,
		).PlaceholderFormat(sq.Dollar).
		From(repository.TransactionTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
				repository.UserTableName,
				repository.UserTableName,
				repository.UserIdColumn,
				repository.TransactionTableName,
				repository.TransactionIdSenderColumn,
			),
		).Where(sq.Eq{
			fmt.Sprintf("%s.%s", 
			repository.TransactionTableName,
			repository.TransactionIdReceiverColumn,
		): receiverID})
	return r.getWithRequest(ctx, tx, builder)
}

func (r *repo) GetTransactionsBySenderID(ctx context.Context, tx pgx.Tx, senderID int64) ([]repoModel.UserTransaction, error) {
	builder := sq.Select(
			repository.UserUsernameColumn,
			repository.TransactionAmountColumn,
		).PlaceholderFormat(sq.Dollar).
		From(repository.TransactionTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
				repository.UserTableName,
				repository.UserTableName,
				repository.UserIdColumn,
				repository.TransactionTableName,
				repository.TransactionIdReceiverColumn,
			),
		).Where(sq.Eq{
			fmt.Sprintf("%s.%s", 
			repository.TransactionTableName,
			repository.TransactionIdSenderColumn,
		): senderID})
	return r.getWithRequest(ctx, tx, builder)
}