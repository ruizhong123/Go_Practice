package tools

import (
	"time"
)

type mockDB struct{}


var mockloginDetails = map[string]*LoginDetails{
	"alex":	{
		Authorization: "123abc",
		Username:      "alex",
	},
	"bob":	{
		Authorization: "456def",
		Username:      "bob",
	},
	"marie":	{
		Authorization: "789ghi",
		Username:      "marie",
	},
}	

vart mockCoinDetails = map[string]CoinDetails{
	"alex":	{
		Coin:     1000,
		Username: "alex",
	},
	"bob":	{
		Coin:     500,
		Username: "bob",
	},
	"marie":	{
		Coin:     750,
		Username: "marie",
	},
}



func (db *mockDB) GetLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second) // 模擬資料庫延遲
	var clientData = LoginDetails{}
	
	clientData,ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}