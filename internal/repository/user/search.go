package user

import (
	"context"
	"errors"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) FindUserByID(ctx context.Context, tx pgx.Tx, userID int64) (*model.User, error) {
	builder := sq.Select(
		repository.UserIdColumn,
		repository.UserUsernameColumn,
		repository.UserPasswordColumn,
		repository.UserCoinsColumn,
	).PlaceholderFormat(sq.Dollar).
	From(repository.UserTableName).
	Where(sq.Eq{repository.UserIdColumn: userID})
	
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errs.NewDBError(err.Error(), err)
	}
	newUser := &model.User{}
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)	
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}
	err = row.Scan(
		&newUser.ID,
		&newUser.Creds.Username,
		&newUser.Creds.PasswordHash,
		&newUser.Coins,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewDBError(err.Error(), err)
	}
	return newUser, nil
}

func (r *repo) FindUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*model.User, error) {
	builder := sq.Select(
		repository.UserIdColumn,
		repository.UserUsernameColumn,
		repository.UserPasswordColumn,
		repository.UserCoinsColumn,	
	).PlaceholderFormat(sq.Dollar).
	From(repository.UserTableName).
	Where(sq.Eq{repository.UserUsernameColumn: username})
	
	query, args, err := builder.ToSql()
	if err != nil {
		errs.NewDBError(err.Error(), err)
	}
	newUser := &model.User{}
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}
	err = row.Scan(
		&newUser.ID,
		&newUser.Creds.Username,
		&newUser.Creds.PasswordHash,
		&newUser.Coins,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewDBError(err.Error(), err)
	}
	return newUser, nil
}