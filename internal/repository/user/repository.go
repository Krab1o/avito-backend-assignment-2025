package user

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/jackc/pgx/v5"
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

func (r *repo) WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
    tx, err := r.db.Begin(ctx)
    if err != nil {
        return errs.NewDBError(err.Error(), err)
    }
    // Defer rollback in case of an error
    defer func() {
        if err != nil {
            tx.Rollback(ctx)
        }
    }()
    // Execute the transactional function
    err = fn(tx)
    if err != nil {
        return err
    }
    // Commit transaction if everything is successful
    return tx.Commit(ctx)
}
