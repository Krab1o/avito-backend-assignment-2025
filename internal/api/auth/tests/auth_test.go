package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/mocks"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	type authServiceMockFunc func (c context.Context, mc *minimock.Controller) service.AuthService

	type args struct {
		c *gin.Context
	}

	type resp struct {
		status	int
		body	string
	}

	var (
		mc = minimock.NewController(t)

		username = gofakeit.Name()
		password = gofakeit.Word()
		reqBody = fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)
		
		serviceCreds = &model.UserCreds{
			Username: username,
			Password: password,
		}

		token = "mocked_token"
	)

	defer mc.Cleanup(mc.Finish)

	tests := []struct{
		name			string
		args			args
		want			resp
		authServiceMock	authServiceMockFunc
	}{
		{
			name: "success case",
			want: resp{
				status: http.StatusOK,
				body: fmt.Sprintf(`{"%s":"%s"}`, api.FieldToken, token),
			},
			authServiceMock: func (ctx context.Context, mc *minimock.Controller) service.AuthService {
				mock := mocks.NewAuthServiceMock(mc)
				mock.AuthMock.Expect(ctx, serviceCreds).Return(token, nil)
				return mock
			},
		},
		{
			name: "internal error case",
			want: resp{
				status: http.StatusUnauthorized,
				body: fmt.Sprintf(`{"%s": "%s"}`, api.FieldError, service.MessageInternalError),
			},
			authServiceMock: func (ctx context.Context, mc *minimock.Controller) service.AuthService {
				mock := mocks.NewAuthServiceMock(mc)
				mock.AuthMock.Expect(ctx, serviceCreds).Return("", errs.NewUnauthorizedError(service.MessageInternalError, nil))
				return mock
			},
		},
		{
			name: "wrong password case",
			want: resp{
				status: http.StatusUnauthorized,
				body: fmt.Sprintf(`{"%s": "%s"}`, api.FieldError, service.MessageWrongPassword),
			},
			authServiceMock: func (ctx context.Context, mc *minimock.Controller) service.AuthService {
				mock := mocks.NewAuthServiceMock(mc)
				mock.AuthMock.Expect(ctx, serviceCreds).Return("", errs.NewUnauthorizedError(service.MessageWrongPassword, nil))
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func (t* testing.T) {
			//Mock request created
			req := httptest.NewRequest(http.MethodPost, api.AuthPath, strings.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			//Mock response recorder
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			ctx := req.Context()

			authServiceMock := tt.authServiceMock(ctx, mc)
			handler := auth.NewHandler(authServiceMock)

			handler.Auth(c)

			//Check for correct written response
			require.Equal(t, tt.want.status, w.Code)
			require.JSONEq(t, tt.want.body, w.Body.String())
		})
	}
}