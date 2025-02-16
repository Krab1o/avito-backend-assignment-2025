package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/user"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/info/model"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/mocks"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/shared"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	type infoServiceMockFunc func (c context.Context, mc *minimock.Controller) service.InfoService

	type args struct {
		c *gin.Context
	}

	type resp struct {
		status	int
		body	string
	}

	var (
		mc = minimock.NewController(t)

		mockedToken = "MOCKED_TOKEN"
		mockedUserID = int64(1234) // Assume this is extracted from the token

		mockedInfo = &model.Info{
			Coins: gofakeit.Number(0, 1000),
			Inventory: model.Inventory(map[string]int{
				"Cup": gofakeit.Number(1, 5),
				"Socks": gofakeit.Number(1, 5),
			}),
			CoinHistory: model.CoinHistory{
				Received: []model.Received{
					{
						FromUser: gofakeit.Name(),
						Amount: gofakeit.Number(5, 25),
					},
					{
						FromUser: gofakeit.Name(),
						Amount: gofakeit.Number(5, 25),
					},
				},
				Sent: []model.Sent{
					{
						ToUser: gofakeit.Name(),
						Amount: gofakeit.Number(5, 25),
					},
					{
						ToUser: gofakeit.Name(),
						Amount: gofakeit.Number(5, 25),
					},
				},
			},
		}
		bodySchema = fmt.Sprintf(`{
			"coins": %d,
			"inventory": [
				{
					"type": "Cup",
					"quantity": %d
				},
				{
					"type": "Socks",
					"quantity": %d
				}
			],
			"coinHistory": {
				"received": [
					{
						"fromUser": "%s",
						"amount": %d
					},
					{
						"fromUser": "%s",
						"amount": %d
					}
				],
				"sent": [
					{
						"toUser": "%s",
						"amount": %d
					},
					{
						"toUser": "%s",
						"amount": %d
					}
				]
			}
		}`, mockedInfo.Coins,
		mockedInfo.Inventory["Cup"], mockedInfo.Inventory["Socks"],
		mockedInfo.CoinHistory.Received[0].FromUser, mockedInfo.CoinHistory.Received[0].Amount,
		mockedInfo.CoinHistory.Received[1].FromUser, mockedInfo.CoinHistory.Received[1].Amount,
		mockedInfo.CoinHistory.Sent[0].ToUser, mockedInfo.CoinHistory.Sent[0].Amount,
		mockedInfo.CoinHistory.Sent[1].ToUser, mockedInfo.CoinHistory.Sent[1].Amount)
	)

	defer mc.Cleanup(mc.Finish)

	tests := []struct{
		name			string
		args			args
		want			resp
		infoServiceMock	infoServiceMockFunc
	}{
		{
			name: "success case",
			want: resp{
				status: http.StatusOK,
				body: bodySchema,
			},
			infoServiceMock: func (ctx context.Context, mc *minimock.Controller) service.InfoService {
				mock := mocks.NewInfoServiceMock(mc)
				mock.InfoMock.Expect(ctx, mockedUserID).Return(mockedInfo, nil)
				return mock
			},
		},
		{
			name: "internal error case",
			want: resp{
				status: http.StatusInternalServerError,
				body: fmt.Sprintf(`{"%s": "%s"}`, api.FieldError, service.MessageInternalError),
			},
			infoServiceMock: func (ctx context.Context, mc *minimock.Controller) service.InfoService {
				mock := mocks.NewInfoServiceMock(mc)
				mock.InfoMock.Expect(ctx, mockedUserID).Return(nil, errs.NewServiceError(service.MessageInternalError, nil))
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func (t* testing.T) {
			//Mock response created
			req := httptest.NewRequest(http.MethodPost, api.AuthPath, nil)
			req.Header.Set("Authorization", "Bearer "+mockedToken)

			//Mock response recorder
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Set(shared.UserIDJsonName, mockedUserID)
			ctx := req.Context()

			infoServiceMock := tt.infoServiceMock(ctx, mc)

			handler := user.NewHandler(infoServiceMock)

			handler.Info(c)

			//Check for correct written response
			expectedBody := tt.want.body
			require.Equal(t, tt.want.status, w.Code)
			var expectedJSON map[string]interface{}
			if err := json.Unmarshal([]byte(expectedBody), &expectedJSON); err != nil {
				t.Fatal("Error unmarshalling expected response:", err)
			}
			indentedExpectedBody, err := json.MarshalIndent(expectedJSON, "", "  ")
			if err != nil {
				t.Fatal("Error formatting expected response:", err)
			}
			require.JSONEq(t, string(indentedExpectedBody), w.Body.String())
		})
	}
}