package auth

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
)

type serv struct {
	userRepository 	repository.UserRepository
	jwtConfig		config.JWTConfig
}

func NewHandler(
		usr repository.UserRepository,
		jwt config.JWTConfig,
	) service.AuthService {
	return &serv{
		userRepository: usr,
		jwtConfig: jwt,
	}
}
