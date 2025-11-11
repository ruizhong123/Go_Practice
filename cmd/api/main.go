package main

import (
	"fmt"
	"github.com/avukadin/goapi/internal/handlers"
	"github.com/go-chi/chi"          // 網路開發套件
	log "github.com/sirupsen/logrus" // 日誌套件
	"net/http"
)

func main() {
	// 啟用日誌呼叫位置
	log.SetReportCaller(true)
	// 引用內部路由器套件啟動路由
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting Go API Server...")
	fmt.Println(
		`  
 /$$$$$$   /$$$$$$         /$$$$$$  /$$$$$$$  /$$$$$$
 /$$__  $$ /$$__  $$       /$$__  $$| $$__  $$|_  $$_/
| $$  \__/| $$  \ $$      | $$  \ $$| $$  \ $$  | $$  
| $$ /$$$$| $$  | $$      | $$$$$$$$| $$$$$$$/  | $$  
| $$|_  $$| $$  | $$      | $$__  $$| $$____/   | $$  
| $$  \ $$| $$  | $$      | $$  | $$| $$        | $$  
|  $$$$$$/|  $$$$$$/      | $$  | $$| $$       /$$$$$$
 \______/  \______/       |__/  |__/|__/      |______/`)

	// LListenAndServe 在正常運行的情況下是部會回傳，
	// 除非遇到錯誤停止運行時，才會回傳錯誤

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)

	}

}
