package auth

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewHandler(usr repository.UserRepository) service.AuthService {
	return &serv{
		userRepository: usr,
	}
}

func (s *serv) Auth(context.Context) error {
	return nil
}
