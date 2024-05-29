package server

import (
	"crud/internal/api"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New() *http.Server {
	r := chi.NewRouter()

	// Создаем экземпляр структуры Handler
	handler := &api.Handler{}

	// Регистрируем обработчики маршрутов
	r.Get("/ping", handler.PingHandler)
	r.Post("/insert_post", handler.InsertPostHandler)

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
