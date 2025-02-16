package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	repoMocks "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/mocks"
	repoTx "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	repoUser "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	txsServ "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction"
	servUser "github.com/Krab1o/avito-backend-assignment-2025/internal/service/transaction/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestBuying(t *testing.T) {
	type mockBehavior func(ctx context.Context, mc *minimock.Controller) (
		userRepo	  	*repoMocks.UserRepositoryMock,
		transactionRepo *repoMocks.TransactionRepositoryMock,
	)

	var (
		mc = minimock.NewController(t)
		mockedSenderID = int64(12)
		mockedSenderCoins = gofakeit.Number(500, 1000)
		mockedSenderName = gofakeit.Username()
		mockedSenderHashedPassword = gofakeit.Password(true, true, true, false, false, 10)

		mockedReceiverID = int64(15)
		mockedReceiverCoins = gofakeit.Number(500, 1000)
		mockedReceiverName = gofakeit.Username()
		mockedReceiverHashedPassword = gofakeit.Password(true, true, true, false, false, 10)

		mockedCoins = gofakeit.Number(10, 50)
		dbErrorMessage = "DB error"
		dbError = errs.NewDBError(dbErrorMessage, nil)

		mockedTx = pgx.Tx(nil)

		tx = &servUser.Transaction{
			FromUser: mockedSenderID,
			ToUser: mockedReceiverName,
			Amount: mockedCoins,
		}


		userReceiver = &repoUser.User{
			ID: mockedReceiverID,
			Creds: repoUser.UserCreds{
				Username: mockedReceiverName,
				PasswordHash: mockedReceiverHashedPassword,
			},
			Coins: mockedReceiverCoins,
		}
		userSender = &repoUser.User{
			ID: mockedSenderID,
			Creds: repoUser.UserCreds{
				Username: mockedSenderName,
				PasswordHash: mockedSenderHashedPassword,
			},
			Coins: mockedSenderCoins,
		}

		txRepo = &repoTx.Transaction{
			SenderID: userSender.ID,
			ReceiverID: userReceiver.ID,
			Amount: tx.Amount,
		}
	)
	
	tests := []struct {
		name         string
		id           int64
		transaction  *servUser.Transaction
		err			 error
		mockBehavior mockBehavior
	}{
		{
			name: "Success - coin sended",
			transaction: tx,
			err: nil,
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(userReceiver, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, tx.FromUser).
					Return(userSender, nil)
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})
				repoUserMock.SubtractCoinsMock.Expect(ctx, mockedTx, userSender, tx.Amount).
					Return(nil)
				repoUserMock.AddCoinsMock.Expect(ctx, mockedTx, userReceiver, tx.Amount).
					Return(nil)
				
				repoTxsMock.CreateTransactionMock.Expect(ctx, mockedTx, txRepo).
					Return(nil)
					
				return repoUserMock, repoTxsMock
			},
		},
		{
			name: "Failure - db error on receiver",
			transaction: tx,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(nil, dbError)
					
				return repoUserMock, repoTxsMock
			},
		},
		{
			name: "Failure - user not found",
			transaction: tx,
			err: errs.NewNotFoundError(service.MessageReceiverNotFound, nil),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(nil, nil)
					
				return repoUserMock, repoTxsMock
			},
		},
		
		{
			name: "Failure - db error on sender",
			transaction: tx,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(userReceiver, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, tx.FromUser).
					Return(nil, dbError)
				return repoUserMock, repoTxsMock
			},
		},
		{
			name: "Failure - receiver is nil",
			transaction: tx,
			err: errs.NewNotFoundError(service.MessageReceiverNotFound, nil),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(nil, nil)
					
				return repoUserMock, repoTxsMock
			},
		},
		{
			name: "Failure - not enough coins",
			transaction: tx,
			err: errs.NewSemanticError(service.MessageNotEnoughCoins, nil),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
			
				userSenderLocal := &repoUser.User{
					ID: mockedSenderID,
					Creds: repoUser.UserCreds{
						Username: mockedSenderName,
						PasswordHash: mockedSenderHashedPassword,
					},
					Coins: 5,
				}

				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(userReceiver, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, tx.FromUser).
					Return(userSenderLocal, nil)
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})
					
				return repoUserMock, repoTxsMock
			},
		},
		{
			name: "Failure - db on subtracting",
			transaction: tx,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(userReceiver, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, tx.FromUser).
					Return(userSender, nil)
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})
				repoUserMock.SubtractCoinsMock.Expect(ctx, mockedTx, userSender, tx.Amount).
					Return(dbError)
					
				return repoUserMock, repoTxsMock
			},
		},
		{
			name: "Failure - db on adding",
			transaction: tx,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(userReceiver, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, tx.FromUser).
					Return(userSender, nil)
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})
				repoUserMock.SubtractCoinsMock.Expect(ctx, mockedTx, userSender, tx.Amount).
					Return(nil)
				repoUserMock.AddCoinsMock.Expect(ctx, mockedTx, userReceiver, tx.Amount).
					Return(dbError)
					
				return repoUserMock, repoTxsMock
			},
		},
		{
			name: "Failure - db on creating transaction",
			transaction: tx,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
			) {
				// Mock repository user
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoUserMock.GetUserByUsernameMock.Expect(ctx, nil, tx.ToUser).
					Return(userReceiver, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, tx.FromUser).
					Return(userSender, nil)
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})
				repoUserMock.SubtractCoinsMock.Expect(ctx, mockedTx, userSender, tx.Amount).
					Return(nil)
				repoUserMock.AddCoinsMock.Expect(ctx, mockedTx, userReceiver, tx.Amount).
					Return(nil)
				repoTxsMock.CreateTransactionMock.Expect(ctx, mockedTx, txRepo).
					Return(dbError)
					
				return repoUserMock, repoTxsMock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creating default gin request context without additions
			// TODO: we can move gin request context to test arguments
			// but now for the sake of simplicity (and also we won't test it function
			// with another context) we will leave it here
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest(http.MethodPost, api.AuthPath, nil)
			c.Request = req
			ctx := c.Request.Context()

			userRepoMock, txsRepoMock := tt.mockBehavior(ctx, mc)
			

			// Create service instance with mocked dependencies
			infoService := txsServ.NewService(txsRepoMock, userRepoMock)
			// buyingService := buying.(userRepoMock, authHelperMock, jwtConfigMock)

			// Call Info
			err := infoService.SendCoin(ctx, tt.transaction)

			require.Equal(t, tt.err, err)
		})
	}
}