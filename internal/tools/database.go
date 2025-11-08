package tools

import (
	log "github.com/sirupsen/logrus"
)

// 決定 API 傳送資料類型
type LoginDetails struct {
	Authorization string
	Username      string
}

type CoinDetail struct {
	Coin     int64
	Username string
}

type DatabaseInterface interface {
	// 取得使用者登入資訊
	GetLoginDetails(username string) *LoginDetails
	// 取得使用者餘額
	GetCoinDetails(username string) *CoinDetail
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB()
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error("database setup error:", err)
		return nil, err
	}
	return &database, nil
}
