package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	repoInv "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	repoMocks "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/mocks"
	repoTxs "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/transaction/model"
	repoUser "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	infoServ "github.com/Krab1o/avito-backend-assignment-2025/internal/service/info"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestBuying(t *testing.T) {
	type mockBehavior func(ctx context.Context, mc *minimock.Controller) (
		userRepo	  	*repoMocks.UserRepositoryMock,
		transactionRepo *repoMocks.TransactionRepositoryMock,
		inventoryRepo 	*repoMocks.InventoryRepositoryMock,
	)

	var (
		mc = minimock.NewController(t)
		mockedUserID = int64(1234)
		mockedUsername = gofakeit.Name()
		mockedHashedPassword = "hashedPassword"
		mockedCoins = gofakeit.Number(100, 1000)

		dbErrorMessage = "DB error"
		dbError = errs.NewDBError(dbErrorMessage, nil)

		sendings = []repoTxs.UserTransaction{
			{
				Username: gofakeit.Name(),
				Amount: gofakeit.Number(5, 200),
			},
			{
				Username: gofakeit.Name(),
				Amount: gofakeit.Number(5, 200),
			},
			{
				Username: gofakeit.Name(),
				Amount: gofakeit.Number(5, 200),
			},
		}

		receivings = []repoTxs.UserTransaction{
			{
				Username: gofakeit.Name(),
				Amount: gofakeit.Number(5, 200),
			},
			{
				Username: gofakeit.Name(),
				Amount: gofakeit.Number(5, 200),
			},
			{
				Username: gofakeit.Name(),
				Amount: gofakeit.Number(5, 200),
			},
		}

		user = &repoUser.User{
			ID: mockedUserID,
			Creds: repoUser.UserCreds{
				Username: mockedUsername,
				PasswordHash: mockedHashedPassword,
			},
			Coins: mockedCoins,
		}

		inventory = repoInv.Inventory{
			"car": 5,
			"cup": 3,
			"funny_merch": 10,
		}

		info = &model.Info{
			Coins: mockedCoins,
			Inventory: model.Inventory{
				"car": 5,
				"cup": 3,
				"funny_merch": 10,
			},
			CoinHistory: model.CoinHistory{
				Received: []model.Received{
					{
						FromUser: receivings[0].Username,
						Amount: receivings[0].Amount,
					},
					{
						FromUser: receivings[1].Username,
						Amount: receivings[1].Amount,
					},
					{
						FromUser: receivings[2].Username,
						Amount: receivings[2].Amount,
					},
				},
				Sent: []model.Sent{
					{
						ToUser: sendings[0].Username,
						Amount: sendings[0].Amount,
					},
					{
						ToUser: sendings[1].Username,
						Amount: sendings[1].Amount,
					},
					{
						ToUser: sendings[2].Username,
						Amount: sendings[2].Amount,
					},
				},
			},
		}
	)
	
	tests := []struct {
		name         string
		id           int64
		info		 *model.Info
		err			 error
		mockBehavior mockBehavior
	}{
		{
			name: "Success - get all info",
			id: mockedUserID,
			info: info,
			err: nil,
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
				inventoryRepo *repoMocks.InventoryRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoTxsMock.GetTransactionsBySenderIDMock.Expect(ctx, nil, mockedUserID).
					Return(sendings, nil)
				repoTxsMock.GetTransactionsByReceiverIDMock.Expect(ctx, nil, mockedUserID).
					Return(receivings, nil)
				repoInventoryMock.GetInventoryByIDMock.Expect(ctx, nil, mockedUserID).
					Return(inventory, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, mockedUserID).
					Return(user, nil)
					
				return repoUserMock, repoTxsMock, repoInventoryMock 
			},
		},
		{
			name: "Failure - Cannot get transaction by sender",
			id: mockedUserID,
			info: nil,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
				inventoryRepo *repoMocks.InventoryRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoTxsMock.GetTransactionsBySenderIDMock.Expect(ctx, nil, mockedUserID).
					Return(nil, dbError)
					
				return repoUserMock, repoTxsMock, repoInventoryMock 
			},
		},
		{
			name: "Failure - Cannot get transaction by receiver",
			id: mockedUserID,
			info: nil,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
				inventoryRepo *repoMocks.InventoryRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoTxsMock.GetTransactionsBySenderIDMock.Expect(ctx, nil, mockedUserID).
					Return(sendings, nil)
				repoTxsMock.GetTransactionsByReceiverIDMock.Expect(ctx, nil, mockedUserID).
					Return(nil, dbError)
					
				return repoUserMock, repoTxsMock, repoInventoryMock 
			},
		},
		{
			name: "Failure - Cannot get inventory by ID",
			id: mockedUserID,
			info: nil,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
				inventoryRepo *repoMocks.InventoryRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoTxsMock.GetTransactionsBySenderIDMock.Expect(ctx, nil, mockedUserID).
					Return(sendings, nil)
				repoTxsMock.GetTransactionsByReceiverIDMock.Expect(ctx, nil, mockedUserID).
					Return(receivings, nil)
				repoInventoryMock.GetInventoryByIDMock.Expect(ctx, nil, mockedUserID).
					Return(nil, dbError)
					
				return repoUserMock, repoTxsMock, repoInventoryMock 
			},
		},
		{
			name: "Failure - Cannot get user by ID",
			id: mockedUserID,
			info: nil,
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				userRepo	  *repoMocks.UserRepositoryMock,
				transactionRepo *repoMocks.TransactionRepositoryMock,
				inventoryRepo *repoMocks.InventoryRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoTxsMock := repoMocks.NewTransactionRepositoryMock(mc)
				
				repoTxsMock.GetTransactionsBySenderIDMock.Expect(ctx, nil, mockedUserID).
					Return(sendings, nil)
				repoTxsMock.GetTransactionsByReceiverIDMock.Expect(ctx, nil, mockedUserID).
					Return(receivings, nil)
				repoInventoryMock.GetInventoryByIDMock.Expect(ctx, nil, mockedUserID).
					Return(inventory, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, mockedUserID).
					Return(nil, dbError)
					
				return repoUserMock, repoTxsMock, repoInventoryMock 
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

			userRepoMock, txsRepoMock, inventoryRepoMock := tt.mockBehavior(ctx, mc)
			
			// Create service instance with mocked dependencies
			infoService := infoServ.NewService(userRepoMock, txsRepoMock, inventoryRepoMock)
			// buyingService := buying.(userRepoMock, authHelperMock, jwtConfigMock)

			// Call Info
			info, err := infoService.Info(ctx, tt.id)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.info, info)
		})
	}
}