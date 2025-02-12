package buying

import "github.com/Krab1o/avito-backend-assignment-2025/internal/service"

const (
	errorNoParamSpecified = "No parameter specified"
	errorInternalServerError = "Internal server error"
)

type Handler struct {
	service service.BuyingService
}

func NewHandler(buyingService service.BuyingService) *Handler {
	return &Handler{
		service: buyingService,
	}
}
