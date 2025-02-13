package buying

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	buyingModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
)


type serv struct {
	inventoryRepo repository.InventoryRepository
}

func NewService(inv repository.InventoryRepository) service.BuyingService {
	return &serv{
		inventoryRepo: inv,
	}
}

func (s *serv) Buy(ctx context.Context, model *buyingModel.Buying) error {
	return nil
}
