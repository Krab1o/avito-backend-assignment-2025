package buying

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/jackc/pgx/v5"
)

const defaultBuyQuantity = 1

//TODO: COMPLETE THROUGHOUT PROPER HANDLING ERRORS FROM LAYER TO LAYER
func (s *serv) Buy(ctx context.Context, newBuying *model.Buying) error {
	merchRepo, err := s.merchRepo.GetItem(ctx, nil, newBuying.Name)
	if err != nil {
		return errs.NewServiceError(service.MessageInternalError, err)
	}
	if merchRepo == nil {
		return errs.NewNotFoundError(service.MessageMerchNotFound, err)
	}
	userRepo, err := s.userRepo.FindUserByID(ctx, nil, newBuying.BuyerID)
	if err != nil {
		return err
	}
	buyingRepo := converter.NewBuying(
		merchRepo.ID,
		userRepo.ID,
		defaultBuyQuantity,
	)
	s.userRepo.WithTransaction(ctx, func (tx pgx.Tx) error {
		var err error
		if userRepo.Coins < merchRepo.Price {
			return errs.NewSemanticError(service.MessageNotEnoughCoins, err)
		}
		err = s.userRepo.SubtractCoins(ctx, tx, userRepo, merchRepo.Price)
		if err != nil {
			return errs.NewServiceError(service.MessageInternalError, err)
		}
		err = s.inventoryRepo.BuyItem(ctx, tx, buyingRepo) 
		if err != nil {
			return errs.NewServiceError(service.MessageInternalError, err)
		}
		return nil
	})
	return nil
}