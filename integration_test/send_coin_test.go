package integrationtest

import (
	"net/http"
	"testing"
)

func TestSending(t *testing.T) {
	//TODO: automate counting
	tests := []struct {
		name 		 string
		coins		 int
		err			 error
	}{
		{
			name : "Success test",
			coins: 10,
			err: nil,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			baseURL := setupServer(t)
			userSender := createUser(t, client, baseURL)
			userReceiver := createUser(t, client, baseURL)
			sendCoin(t, client, baseURL, userSender, userReceiver, 10)
			getInfo(t, client, baseURL, userSender)
			getInfo(t, client, baseURL, userReceiver)
		})
	}
}