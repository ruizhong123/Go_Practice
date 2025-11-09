package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123abc",
		Username:  "alex",
	},
	"bob": {
		AuthToken: "456def",
		Username:  "bob",
	},
	"marie": {
		AuthToken: "789ghi",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    1000,
		Username: "alex",
	},
	"bob": {
		Coins:    500,
		Username: "bob",
	},
	"marie": {
		Coins:    750,
		Username: "marie",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second) // 模擬資料庫延遲
	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (db *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(1 * time.Second) // 模擬資料庫延遲
	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}
