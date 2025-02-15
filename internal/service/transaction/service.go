package transaction

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
)

type serv struct {
	transactionRepo repository.TransactionRepository
	userRepo 		repository.UserRepository
}

func NewService(
		tr repository.TransactionRepository,
		usr repository.UserRepository,
	) service.TransactionService {
	return &serv{
		transactionRepo: tr,
		userRepo:		 usr,
	}
}
