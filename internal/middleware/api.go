package middleware

import (
	"error"
	"net/http"

	"github.com/avukadin/goapi.git/api"
	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = error.New("unauthorized access")

// 使用 Authorization 處理使用著 授權
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		var err error

		// 如果相位
		if username == "" || token == "" {
			log.Error("unauthorizedError")
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.LoginDetails
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails := (*database).GetLoginDetails(username)

		if (loginDetails == nil) || (token != (*loginDetails).AuthToken) {
			log.Error("unauthorizedError")
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
