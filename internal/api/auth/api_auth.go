package auth

import "github.com/Krab1o/avito-backend-assignment-2025/internal/service"

type Handler struct {
	authService service.AuthService
}

func NewHandler(authService service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}