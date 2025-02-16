package auth

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	authhelper "github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/auth_helper"
)

type serv struct {
	userRepository 	repository.UserRepository
	authHelper		authhelper.AuthHelper
	jwtConfig		config.JWTConfig
}

func NewService(
		usr repository.UserRepository,
		helper authhelper.AuthHelper,
		jwt config.JWTConfig,
	) service.AuthService {
	return &serv{
		userRepository: usr,
		authHelper: helper,
		jwtConfig: jwt,
	}
}
