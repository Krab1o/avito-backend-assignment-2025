package transaction

import (
	"context"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/jackc/pgx/v5"
)

func (s *serv) SendCoin(ctx context.Context, newTransaction *model.Transaction) error {	
	//TODO: better do using goroutines (parallel reading)
	receiver, err := s.userRepo.GetUserByUsername(ctx, nil, newTransaction.ToUser)
	if err != nil {
		return errs.NewServiceError(service.MessageInternalError, err)
	}
	if receiver == nil {
		return errs.NewNotFoundError(service.MessageReceiverNotFound, err)
	}
	// No need to check sender because of assumption that it sent valid token
	sender, err := s.userRepo.GetUserByID(ctx, nil, newTransaction.FromUser)
	if err != nil {
		return errs.NewServiceError(service.MessageInternalError, err)
	}
	return s.userRepo.WithTransaction(ctx, func(tx pgx.Tx) error {		
		//Business logic check
		if receiver == nil || sender == nil {
			return errs.NewSemanticError(service.MessageUserNotFound, err)
		}
		if sender.Coins < newTransaction.Amount {
			return errs.NewSemanticError(service.MessageNotEnoughCoins, err)
		}
		if sender.ID == receiver.ID {
			return errs.NewSemanticError(service.MessageSelfSending, err)
		}
		//Continue transaction
		err = s.userRepo.SubtractCoins(ctx, tx, sender, newTransaction.Amount)
		if err != nil {
			return errs.NewServiceError(service.MessageInternalError, err)
		}
		err = s.userRepo.AddCoins(ctx, tx, receiver, newTransaction.Amount)
		if err != nil {
			return errs.NewServiceError(service.MessageInternalError, err)
		}
		newTransactionRepo := converter.TransactionServiceToRepo(newTransaction, receiver.ID)
		err = s.transactionRepo.CreateTransaction(ctx, tx, newTransactionRepo)
		if err != nil {
			return errs.NewServiceError(service.MessageInternalError, err)
		}
		return nil
	})
}