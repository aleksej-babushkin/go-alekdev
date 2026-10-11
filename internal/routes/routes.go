package routes

import (
	"app/internal/handlers"
	"app/internal/services"
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	orderService := services.NewOrderService()
	orderHandler := handlers.NewOrderHandler(orderService)

	mux.HandleFunc("POST /api/v1/orders", orderHandler.Create)

	return mux
}
