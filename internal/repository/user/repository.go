package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) repository.UserRepository {
	return &repo{
		db: pool,
	}
}

func (r *repo) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	stmt := fmt.Sprintf(
		`INSERT INTO %s (%s, %s, %s) 
		VALUES ($1, $2, $3)
		RETURNING %s`,
		repository.UserTableName,
		repository.UserUsernameColumn,
		repository.UserPasswordColumn,
		repository.UserCoinsColumn,
		repository.UserIdColumn,
	)
	var id int64
	err := r.db.QueryRow(ctx, stmt, 
		user.Creds.Username, 
		user.Creds.PasswordHash, 
		user.Coins,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) FindUserCreds(ctx context.Context, creds *model.UserCreds) (*model.UserCreds, error) {
	stmt := fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s = $1", 
		repository.UserUsernameColumn,
		repository.UserPasswordColumn,
		repository.UserTableName,
		repository.UserUsernameColumn,
	)
	user := &model.UserCreds{}
	err := r.db.QueryRow(ctx, stmt, creds.Username).Scan(
		user.Username,
		user.PasswordHash,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

//TODO: implement method
func (r *repo) Info(ctx context.Context) (*model.Info, error) {
	return nil, nil
}