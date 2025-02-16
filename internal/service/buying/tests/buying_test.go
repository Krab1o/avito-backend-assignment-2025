package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	repoInv "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	repoMerch "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/merch/model"
	repoMocks "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/mocks"
	repoUser "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	buyingServ "github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/buying/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestBuying(t *testing.T) {
	type mockBehavior func(ctx context.Context, mc *minimock.Controller) (
		inventoryRepo *repoMocks.InventoryRepositoryMock,
		userRepo	  *repoMocks.UserRepositoryMock,
		merchRepo     *repoMocks.MerchRepositoryMock,
	)

	var (
		mc = minimock.NewController(t)
		mockedMerchId = int64(333)
		mockedMerchTitle = gofakeit.Name()
		mockerMerchPrice = gofakeit.Number(100, 200)
		quantity = 1

		mockedUserID = int64(1234)
		mockedUsername = "testUser"
		mockedHashedPassword = "hashedPass"
		mockedCoins = gofakeit.Number(300, 1000)

		
		dbErrorMessage = "DB error"
		dbError = errs.NewDBError(dbErrorMessage, nil)

		mockedTx = pgx.Tx(nil)


		buying = &repoInv.Buying{
			BuyerID: mockedUserID,
			MerchID: mockedMerchId,
			Quantity: quantity,
		}

		user = &repoUser.User{
			ID: mockedUserID,
			Creds: repoUser.UserCreds{
				Username: mockedUsername,
				PasswordHash: mockedHashedPassword,
			},
			Coins: mockedCoins,
		}

		merch = &repoMerch.Merch{
			ID: mockedMerchId,
			Title: mockedMerchTitle,
			Price: mockerMerchPrice,
		}
	)
	
	tests := []struct {
		name         string
		args         *model.Buying
		err			 error
		mockBehavior mockBehavior
	}{
		{
			name: "Success - Existing user with correct password",
			args: &model.Buying{
				BuyerID: mockedUserID,
				Name: mockedMerchTitle,
			},
			err: nil,
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				inventoryRepo *repoMocks.InventoryRepositoryMock,
				userRepo	  *repoMocks.UserRepositoryMock,
				merchRepo     *repoMocks.MerchRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoMerchMock := repoMocks.NewMerchRepositoryMock(mc)
				
				repoMerchMock.GetItemMock.Expect(ctx, nil, mockedMerchTitle).
					Return(merch, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, mockedUserID).
					Return(user, nil)
				
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})

				repoUserMock.SubtractCoinsMock.Expect(ctx, mockedTx, user, merch.Price).
					Return(nil)

				repoInventoryMock.BuyItemMock.Expect(ctx, mockedTx, buying).
					Return(nil)
					
				return repoInventoryMock, repoUserMock, repoMerchMock
			},
		},
		{
			name: "Failure - failed to get data from merchRepo",
			args: &model.Buying{
				BuyerID: mockedUserID,
				Name: mockedMerchTitle,
			},
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				inventoryRepo *repoMocks.InventoryRepositoryMock,
				userRepo	  *repoMocks.UserRepositoryMock,
				merchRepo     *repoMocks.MerchRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoMerchMock := repoMocks.NewMerchRepositoryMock(mc)
				
				repoMerchMock.GetItemMock.Expect(ctx, nil, mockedMerchTitle).
					Return(nil, dbError)
					
				return repoInventoryMock, repoUserMock, repoMerchMock
			},
		},
		{
			name: "Failure - failed to get data from userRepo",
			args: &model.Buying{
				BuyerID: mockedUserID,
				Name: mockedMerchTitle,
			},
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				inventoryRepo *repoMocks.InventoryRepositoryMock,
				userRepo	  *repoMocks.UserRepositoryMock,
				merchRepo     *repoMocks.MerchRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoMerchMock := repoMocks.NewMerchRepositoryMock(mc)
				
				repoMerchMock.GetItemMock.Expect(ctx, nil, mockedMerchTitle).
					Return(merch, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, mockedUserID).
					Return(nil, dbError)
					
				return repoInventoryMock, repoUserMock, repoMerchMock
			},
		},
		{
			name: "Failure - merch not exists",
			args: &model.Buying{
				BuyerID: mockedUserID,
				Name: mockedMerchTitle,
			},
			err: errs.NewNotFoundError(service.MessageMerchNotFound, nil),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				inventoryRepo *repoMocks.InventoryRepositoryMock,
				userRepo	  *repoMocks.UserRepositoryMock,
				merchRepo     *repoMocks.MerchRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoMerchMock := repoMocks.NewMerchRepositoryMock(mc)
				
				repoMerchMock.GetItemMock.Expect(ctx, nil, mockedMerchTitle).
					Return(nil, nil)
					
				return repoInventoryMock, repoUserMock, repoMerchMock
			},
		},
		{
			name: "Failure - not enough coins",
			args: &model.Buying{
				BuyerID: mockedUserID,
				Name: mockedMerchTitle,
			},
			err: errs.NewSemanticError(service.MessageNotEnoughCoins, nil),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				inventoryRepo *repoMocks.InventoryRepositoryMock,
				userRepo	  *repoMocks.UserRepositoryMock,
				merchRepo     *repoMocks.MerchRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoMerchMock := repoMocks.NewMerchRepositoryMock(mc)
				
				userLocal := &repoUser.User{
					ID: mockedUserID,
					Creds: repoUser.UserCreds{
						Username: mockedUsername,
						PasswordHash: mockedHashedPassword,
					},
					Coins: 500,
				}
		
				merchLocal := &repoMerch.Merch{
					ID: mockedMerchId,
					Title: mockedMerchTitle,
					Price: 1000,
				}

				repoMerchMock.GetItemMock.Expect(ctx, nil, mockedMerchTitle).
					Return(merchLocal, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, mockedUserID).
					Return(userLocal, nil)
				
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})
					
				return repoInventoryMock, repoUserMock, repoMerchMock
			},
		},
		{
			name: "Failure - failed to substract coins in repo",
			args: &model.Buying{
				BuyerID: mockedUserID,
				Name: mockedMerchTitle,
			},
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				inventoryRepo *repoMocks.InventoryRepositoryMock,
				userRepo	  *repoMocks.UserRepositoryMock,
				merchRepo     *repoMocks.MerchRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoMerchMock := repoMocks.NewMerchRepositoryMock(mc)
				
				repoMerchMock.GetItemMock.Expect(ctx, nil, mockedMerchTitle).
					Return(merch, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, mockedUserID).
					Return(user, nil)
				
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})

				repoUserMock.SubtractCoinsMock.Expect(ctx, mockedTx, user, merch.Price).
					Return(dbError)
					
				return repoInventoryMock, repoUserMock, repoMerchMock
			},
		},
		{
			name: "Success - failed to create buying in repo",
			args: &model.Buying{
				BuyerID: mockedUserID,
				Name: mockedMerchTitle,
			},
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				inventoryRepo *repoMocks.InventoryRepositoryMock,
				userRepo	  *repoMocks.UserRepositoryMock,
				merchRepo     *repoMocks.MerchRepositoryMock,
			) {
				// Mock repository user
				repoInventoryMock := repoMocks.NewInventoryRepositoryMock(mc)
				repoUserMock := repoMocks.NewUserRepositoryMock(mc)
				repoMerchMock := repoMocks.NewMerchRepositoryMock(mc)
				
				repoMerchMock.GetItemMock.Expect(ctx, nil, mockedMerchTitle).
					Return(merch, nil)
				repoUserMock.GetUserByIDMock.Expect(ctx, nil, mockedUserID).
					Return(user, nil)
				
				repoUserMock.WithTransactionMock.Set(func (ctx context.Context, fn func(pgx.Tx) error) error {
					return fn(mockedTx)
				})

				repoUserMock.SubtractCoinsMock.Expect(ctx, mockedTx, user, merch.Price).
					Return(nil)

				repoInventoryMock.BuyItemMock.Expect(ctx, mockedTx, buying).
					Return(dbError)
					
				return repoInventoryMock, repoUserMock, repoMerchMock
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

			inventoryRepoMock, userRepoMock, merchRepoMock := tt.mockBehavior(ctx, mc)
			
			// Create service instance with mocked dependencies
			buyingService := buyingServ.NewService(inventoryRepoMock, userRepoMock, merchRepoMock)
			// buyingService := buying.(userRepoMock, authHelperMock, jwtConfigMock)

			// Call Auth
			err := buyingService.Buy(ctx, tt.args)

			require.Equal(t, tt.err, err)
		})
	}
}