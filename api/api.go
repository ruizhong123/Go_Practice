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
	Balance float64
}

// 回傳錯誤訊息
type Error struct {
	Code    int
	Message string
}

func writerError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writerError(w, err.Error(), http.StatusBadRequest)
	}
	internalErrorHandler = func(w http.ResponseWriter) {
		writerError(w, "internal server error", http.StatusInternalServerError)
	}
)
