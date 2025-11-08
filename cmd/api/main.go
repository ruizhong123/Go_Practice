package api

import (
	"fmt"
	"github.com/avukadin/goapi/internal/handlers"
	"github.com/go-chi/chi"          // 網路開發套件
	log "github.com/sirupsen/logrus" // 日誌套件
	"net/http"
)

func main() {
	log.SetReportCaller(true)
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
	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)

	}

}
