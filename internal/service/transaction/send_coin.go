package transaction

import (
	"context"

	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/jackc/pgx/v5"
)

func (s *serv) SendCoin(ctx context.Context, newTransaction *model.Transaction) error {	
	var sender *repoModel.User
	var receiver *repoModel.User
	s.userRepo.WithTransaction(ctx, func(tx pgx.Tx) error {
		var err error
		//TODO: better do using goroutines (parallel reading)
		receiver, err = s.userRepo.FindUserByUsername(ctx, tx, newTransaction.ToUser)
		if err != nil {
			return errs.NewServiceError(service.MessageInternalError, err)
		}
		sender, err = s.userRepo.FindUserByID(ctx, tx, newTransaction.FromUser)
		if err != nil {
			return errs.NewServiceError(service.MessageInternalError, err)
		}
		if receiver == nil || sender == nil {
			return errs.NewSemanticError(service.MessageUserNotFound, err)
		}
		//business logic check
		if sender.Coins < newTransaction.Amount {
			return errs.NewSemanticError(service.MessageNotEnoughCoins, err)
		}
		if sender.ID == receiver.ID {
			return errs.NewSemanticError(service.MessageSelfSending, err)
		}
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
	return nil
}