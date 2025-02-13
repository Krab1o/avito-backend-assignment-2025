package buying

import "github.com/Krab1o/avito-backend-assignment-2025/internal/service"

type Handler struct {
	buyingService service.BuyingService
}

func NewHandler(buyingService service.BuyingService) *Handler {
	return &Handler{
		buyingService: buyingService,
	}
}
