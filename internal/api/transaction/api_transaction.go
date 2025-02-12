package transaction

import "github.com/Krab1o/avito-backend-assignment-2025/internal/service"

const (
	errorNoParamSpecified = "No parameter specified"
	errorInternalServerError = "Internal server error"
	errorBadRequest = "Bad request specified"
)

type Handler struct {
	service service.TransactionService
}

func NewHandler(transactionService service.TransactionService) *Handler {
	return &Handler{
		service: transactionService,
	}
}