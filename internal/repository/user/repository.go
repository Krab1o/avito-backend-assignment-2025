package user

import (
	"context"

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

//TODO: implement method
func (r *repo) CreateUser(ctx context.Context) error {
	return nil
}
//TODO: implement method
func (r *repo) FindUser(ctx context.Context, creds *model.UserCreds) (*model.User, error) {
	return nil, nil
}

func (r *repo) Info(ctx context.Context) (*model.Info, error) {
	return nil, nil
}