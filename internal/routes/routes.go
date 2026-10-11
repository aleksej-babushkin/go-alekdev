package routes

import (
	"app/internal/handlers"
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/orders", handlers.CreateOrder)
	return mux
}
