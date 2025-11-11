package api

import (
	"encoding/json"
	"net/http"
)

////// API 參數與回傳結構 //////

// 建立使用者名稱
type CoinBalanceParam struct {
	Username string
}

// 回傳使用者餘額
type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

// 回傳錯誤訊息//

type Error struct {
	Code    int
	Message string
}

// 撰寫 錯誤格式的函式
func writerError(w http.ResponseWriter, message string, code int) {

	// response 回傳錯誤訊息 : 錯誤代碼與錯誤訊息
	resp := Error{
		Code:    code,
		Message: message,
	}

	// 設定 response 撰寫格式為 JSON
	w.Header().Set("Content-Type", "application/json")

	// 設定 response 狀態碼
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

// 定義兩個錯誤函式變數 : 一個主要處理請求錯誤的會傳，另一個是網路上的錯誤

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writerError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writerError(w, "internal server error", http.StatusInternalServerError)
	}
)
