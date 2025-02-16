package integrationtest

import (
	"net/http"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/user/dto"
)

func TestBuying(t *testing.T) {
	tests := []struct {
		name         string
		info		 *dto.Info
		buyings      map[string]int
		finalCoins	 int
		err			 error
	}{
		//TODO: can be improved by autocounting somehow
		{
			name : "Success test",
			buyings: map[string]int{
				"cup": 5,
				"powerbank": 2,
				"wallet": 3,
			},
			info: &dto.Info{
				Coins: 350,
				Inventory: []dto.Inventory{
					{
						Type: "wallet",
						Quantity: 3,
					},
					{
						Type: "powerbank",
						Quantity: 2,
					},
					{
						Type: "cup",
						Quantity: 5,
					},
				},
				CoinHistory: dto.CoinHistory{
					Received: []dto.Received{},
					Sent: []dto.Sent{},
				},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			baseURL := setupServer(t)
			user := createUser(t, client, baseURL)
			buyMerch(t, client, baseURL, user, tt.buyings)
			//TODO: maybe convert buyings to inventory form so can be validated
			getInfo(t, client, baseURL, user)
		})
	}
}


