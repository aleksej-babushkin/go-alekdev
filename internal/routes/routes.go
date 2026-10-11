package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"app/internal/handlers"
	"app/internal/repositories"
	"app/internal/services"
)

func New(db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	mux.HandleFunc("POST /api/v1/orders", orderHandler.Create)

	return mux
}
