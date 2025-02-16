package integrationtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/user/dto"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/config/env"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/service/auth"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

const (
	authDataUsername = "username"
	authDataPassword = "password"

	sendDataReceiver = "toUser"
	sendDataAmount = "amount"
)

type UserTest struct {
	username 	string
	token 		string
	coins 		int
}

func setupServer(t *testing.T) string {
	config.Load("../.env")
	httpConfig, err := env.NewHTTPConfig()
	if err != nil {
		t.Fatal("Failed to load http config")
	}
	return fmt.Sprintf("http://%s", httpConfig.AccessAddress())
}

func createUser(t *testing.T, client *http.Client, baseURL string) *UserTest {
	var token string
	endpoint := baseURL + api.AuthPath
	var respMapped map[string]string
	authData := map[string]string{
		authDataUsername: gofakeit.Username(),
		authDataPassword: gofakeit.Password(true, true, true, false, false, 10),
	}
	jsonAuthData, err := json.Marshal(authData)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Post(endpoint, "application/json", bytes.NewBuffer(jsonAuthData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(respBody, &respMapped); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, resp.StatusCode, http.StatusOK)
	token = respMapped[api.FieldToken]
	return &UserTest{
		username: authData[authDataUsername],
		token: token,
		coins: auth.StartCoins,
	}
}

func buyMerch(t *testing.T, client *http.Client, baseURL string, user *UserTest, buyings map[string]int) {
	endpoint := baseURL + api.BuyItemPath[:len(api.BuyItemPath) - 5]
	
	for item, count := range buyings {
		for i := 0; i < count; i++ {
			fullURL := endpoint + item
			req, err := http.NewRequest(http.MethodGet, fullURL, nil)
			req.Header.Add("Authorization", "Bearer " + user.token)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			
			require.Equal(t, http.StatusOK, resp.StatusCode)
		}
	} 
}

func sendCoin(
		t *testing.T, 
		client *http.Client, 
		baseURL string, 
		sender *UserTest,
		receiver *UserTest,
		coins int,
	) {
	endpoint := baseURL + api.SendCoinPath
	sendingData := map[string]any{
		sendDataReceiver: receiver.username,
		sendDataAmount: coins,
	}
	jsonSendingData, err := json.Marshal(sendingData)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonSendingData))
	req.Header.Add("Authorization", "Bearer " + sender.token)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode) 
}

func getInfo(t *testing.T, client *http.Client, baseURL string, user *UserTest) *dto.Info{
	endpoint := baseURL + api.InfoPath
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Add("Authorization", "Bearer " + user.token)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	info := &dto.Info{}
	require.Equal(t, http.StatusOK, resp.StatusCode)
	err = json.Unmarshal(body, info)
	if err != nil {
		t.Fatal(err)
	}
	return info
}