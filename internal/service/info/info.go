package info

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
)

func (s *serv) Info(ctx context.Context, id int64) (*model.Info, error) {
	sendTrRepo, err := s.transactionRepo.GetTransactionsBySenderID(ctx, nil, id)
	if err != nil {
		return nil, errs.NewServiceError(service.MessageInternalError, err)
	}
	// log.Println(sendTrRepo)
	receiverTrRepo, err := s.transactionRepo.GetTransactionsByReceiverID(ctx, nil, id)
	if err != nil {
		return nil, errs.NewServiceError(service.MessageInternalError, err)
	}
	// log.Println(receiverTrRepo)
	inventory, err := s.inventoryRepo.GetInventoryByID(ctx, nil, id)
	if err != nil {
		return nil, errs.NewServiceError(service.MessageInternalError, err)
	}
	user, err := s.userRepo.GetUserByID(ctx, nil, id)
	if err != nil {
		return nil, errs.NewServiceError(service.MessageInternalError, err)
	}
	return converter.InfoRepoToService(
		user.Coins,
		sendTrRepo,
		receiverTrRepo,
		inventory,
	), nil
} 