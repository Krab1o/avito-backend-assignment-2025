package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	confMocks "github.com/Krab1o/avito-backend-assignment-2025/internal/config/mocks"
	repoMocks "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/mocks"
	repoModel "github.com/Krab1o/avito-backend-assignment-2025/internal/repository/user/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth"
	authMocks "github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/auth_helper/mocks"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	type mockBehavior func(ctx context.Context, mc *minimock.Controller) (
		*repoMocks.UserRepositoryMock,
		*authMocks.AuthHelperMock,
		*confMocks.JWTConfigMock,
	)

	var (
		mc = minimock.NewController(t)
		mockedUserID = int64(1234)
		mockedUsername = "testUser"
		mockedPassword = "securePass"
		mockedHashedPassword = "hashedPass"
		mockedCoins = gofakeit.Number(0, 1000)
		mockedToken = "mocked_token"
		mockedSecret = []byte("jwt_secret")
		mockedTimeout = gofakeit.Number(10, 60)
		
		dbErrorMessage = "DB error"
		dbError = errs.NewDBError(dbErrorMessage, nil)
	)
	
	tests := []struct {
		name         string
		args         *model.UserCreds
		want		 string
		err			 error
		mockBehavior mockBehavior
	}{
		{
			name: "Success - Existing user with correct password",
			args: &model.UserCreds{
				Username: mockedUsername,
				Password: mockedPassword,
			},
			want: mockedToken,
			err: nil,
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				*repoMocks.UserRepositoryMock,
				*authMocks.AuthHelperMock,
				*confMocks.JWTConfigMock,
			) {
				// Mock repository user
				jwtConfMock := confMocks.NewJWTConfigMock(mc)
				repoMock := repoMocks.NewUserRepositoryMock(mc)
				authHelperMock := authMocks.NewAuthHelperMock(mc)
				repoMock.GetUserByUsernameMock.Expect(ctx, nil, mockedUsername).Return(
					&repoModel.User{
						ID: mockedUserID,
						Creds: repoModel.UserCreds{
							Username: mockedUsername,
							PasswordHash: mockedHashedPassword,
						},
						Coins: mockedCoins,
					}, 
					nil,
				)
				
				jwtConfMock.SecretMock.Return(mockedSecret)
				jwtConfMock.TimeoutMock.Return(mockedTimeout)

				authHelperMock.VerifyPasswordMock.Expect(mockedHashedPassword, mockedPassword).
					Return(true)

				authHelperMock.GenerateJWTMock.Expect(mockedUserID, mockedSecret, mockedTimeout).
					Return(mockedToken, nil)
				// Mock password verification
				return repoMock, authHelperMock, jwtConfMock
			},
		},
		{
			name: "Failure - Existing user with incorrect password",
			args: &model.UserCreds{
				Username: mockedUsername,
				Password: mockedPassword,
			},
			want: "",
			err: errs.NewUnauthorizedError(service.MessageWrongPassword, nil),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				*repoMocks.UserRepositoryMock,
				*authMocks.AuthHelperMock,
				*confMocks.JWTConfigMock,
			) {
				// Mock repository user
				jwtConfMock := confMocks.NewJWTConfigMock(mc)
				repoMock := repoMocks.NewUserRepositoryMock(mc)
				authHelperMock := authMocks.NewAuthHelperMock(mc)
				repoMock.GetUserByUsernameMock.Expect(ctx, nil, mockedUsername).Return(
					&repoModel.User{
						ID: mockedUserID,
						Creds: repoModel.UserCreds{
							Username: mockedUsername,
							PasswordHash: mockedHashedPassword,
						},
						Coins: mockedCoins,
					}, 
					nil,
				)

				authHelperMock.VerifyPasswordMock.Expect(mockedHashedPassword, mockedPassword).
					Return(false)
				
				// Mock password verification
				return repoMock, authHelperMock, jwtConfMock
			},
		},
		{
			name: "Failure - GetUserByUsername returned error",
			args: &model.UserCreds{
				Username: mockedUsername,
				Password: mockedPassword,
			},
			want: "",
			err: errs.NewServiceError(service.MessageInternalError, dbError),
			mockBehavior: func(ctx context.Context, mc *minimock.Controller) (
				*repoMocks.UserRepositoryMock,
				*authMocks.AuthHelperMock,
				*confMocks.JWTConfigMock,
			) {
				// Mock repository user
				jwtConfMock := confMocks.NewJWTConfigMock(mc)
				repoMock := repoMocks.NewUserRepositoryMock(mc)
				authHelperMock := authMocks.NewAuthHelperMock(mc)
				repoMock.GetUserByUsernameMock.Expect(ctx, nil, mockedUsername).Return(
					nil, dbError,
				)
				
				// jwtConfMock.SecretMock.Return(mockedSecret)
				// jwtConfMock.TimeoutMock.Return(mockedTimeout)

				// authHelperMock.VerifyPasswordMock.Expect(mockedHashedPassword, mockedPassword).
				// 	Return(true)

				// authHelperMock.GenerateJWTMock.Expect(mockedUserID, mockedSecret, mockedTimeout).
				// 	Return(mockedToken, nil)
				// Mock password verification
				return repoMock, authHelperMock, jwtConfMock
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

			userRepoMock, authHelperMock, jwtConfigMock := tt.mockBehavior(ctx, mc)

			// Create service instance with mocked dependencies
			authService := auth.NewService(userRepoMock, authHelperMock, jwtConfigMock)

			// Call Auth
			token, err := authService.Auth(ctx, tt.args)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, token)
		})
	}
}