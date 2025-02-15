package user

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) AddCoins(ctx context.Context, tx pgx.Tx, user *model.User, value int) error {
	builder := sq.Update(repository.UserTableName).
		PlaceholderFormat(sq.Dollar).
		Set(repository.UserCoinsColumn, user.Coins + value).
		Where(sq.Eq{repository.UserIdColumn: user.ID})
	query, args, err := builder.ToSql()
	if err != nil {
		return errs.NewDBError(err.Error(), err)
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

func (r *repo) SubtractCoins(ctx context.Context, tx pgx.Tx, user *model.User, value int) error {
	builder := sq.Update(repository.UserTableName).
		PlaceholderFormat(sq.Dollar).
		Set(repository.UserCoinsColumn, user.Coins - value).
		Where(sq.Eq{repository.UserIdColumn: user.ID})
	query, args, err := builder.ToSql()
	if err != nil {
		return errs.NewDBError(err.Error(), err)
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