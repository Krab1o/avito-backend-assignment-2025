package buying

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
)

type buyingService struct {

}

func NewService() service.BuyingService {
	return &buyingService{}
}

func (s *buyingService) Buy(ctx context.Context, itemName string) error {
	return nil
}
