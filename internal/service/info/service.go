package info

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
)

type serv struct {
	userRepo repository.UserRepository
}

func NewService(usr repository.UserRepository) service.InfoService {
	return &serv{
		userRepo: usr,
	}
}