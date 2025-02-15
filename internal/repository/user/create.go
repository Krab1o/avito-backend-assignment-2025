package user

import (
	"context"
	"fmt"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) CreateUser(ctx context.Context, tx pgx.Tx, user *model.User) (int64, error) {
	builder := sq.Insert(repository.UserTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			repository.UserUsernameColumn,
			repository.UserPasswordColumn,
			repository.UserCoinsColumn,
		).Values(
			user.Creds.Username, 
			user.Creds.PasswordHash, 
			user.Coins,
		).Suffix(
			fmt.Sprintf("RETURNING %s", repository.UserIdColumn),
		)
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, errs.NewDBError(err.Error(), err)
	}
	
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, errs.NewDBError(err.Error(), err)
	}
	return id, nil
}