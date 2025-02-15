package info

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
)

type serv struct {
	userRepo repository.UserRepository
	transactionRepo repository.TransactionRepository
	inventoryRepo repository.InventoryRepository
}

func NewService(
		usr repository.UserRepository,
		tr repository.TransactionRepository,
		inv repository.InventoryRepository,
	) service.InfoService {
	return &serv{
		userRepo: usr,
		transactionRepo: tr,
		inventoryRepo: inv,
	}
}