package user

import "github.com/Krab1o/avito-backend-assignment-2025/internal/service"

type Handler struct {
	infoService service.InfoService
}

func NewHandler(infoService service.InfoService) *Handler {
	return &Handler{
		infoService: infoService,
	}
}