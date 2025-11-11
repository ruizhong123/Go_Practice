package handlers

import (
	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	// 先進入授權的中介軟體，在進入取得戶頭餘額
	r.Route("/account", func(router chi.Router) {

		// 使用 Authorization 中介軟體，b

		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})

}
