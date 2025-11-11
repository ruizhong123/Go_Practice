package middleware

import (
	"errors"
	"net/http"

	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("unauthorized access")

// 使用 Authorization 處理使用著 授權
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 根據 request 查詢適當的使用著名稱以及token
		var username string = r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error

		// 如果驗證失敗
		if username == "" || token == "" {
			log.Error("unauthorizedError")
			// 交給 api 文件中 requesthandler 處理錯誤
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface

		database, err = tools.NewDatabase()

		// 網路連線有誤交給requesthandler 處理
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// 根據資料庫取得使用者登入資訊
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil) || (token != (*loginDetails).AuthToken) {
			log.Error("unauthorizedError")
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// *Authorization -> next.ServeHTTP(w, r) -> GetCoinBalance
