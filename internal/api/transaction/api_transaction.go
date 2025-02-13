package transaction

import "github.com/Krab1o/avito-backend-assignment-2025/internal/service"

type Handler struct {
	transactionService service.TransactionService
}

func NewHandler(transactionService service.TransactionService) *Handler {
	return &Handler{
		transactionService: transactionService,
	}
}