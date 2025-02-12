package buying

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	buyingModel "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
)


type buyingService struct {

}

func NewService() service.BuyingService {
	return &buyingService{}
}

func (s *buyingService) Buy(ctx context.Context, model *buyingModel.Buying) error {
	return nil
}
