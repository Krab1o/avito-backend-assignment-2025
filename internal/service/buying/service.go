package buying

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
)


type serv struct {
	inventoryRepo repository.InventoryRepository
	userRepo	  repository.UserRepository
	merchRepo     repository.MerchRepository
}

func NewService(
		inv  repository.InventoryRepository,
		usr  repository.UserRepository,
		mrch repository.MerchRepository,
	) service.BuyingService {
	return &serv{
		inventoryRepo: inv,
		userRepo: usr,
		merchRepo: mrch,
	}
}
