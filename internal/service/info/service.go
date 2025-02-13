package info

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
)

type serv struct {
	userRepo repository.UserRepository
}

func NewService(usr repository.UserRepository) service.InfoService {
	return &serv{
		userRepo: usr,
	}
}

func (s *serv) Info(ctx context.Context) (*model.Info, error) {
	return nil, nil
} 