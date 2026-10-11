package handlers

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"

	"app/internal/dto"
	"app/internal/services"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req dto.CreateOrderRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}

	input := services.CreateOrderInput{
		Name:      req.Name,
		Phone:     req.Phone,
		Message:   req.Message,
		IP:        ip,
		UserAgent: r.UserAgent(),
		SourceURL: r.Referer(),
	}

	id, err := h.service.Create(r.Context(), input)
	if err != nil {
		if errors.Is(err, services.ErrInvalidOrder) {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid order data",
			})
			return
		}

		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "Internal server error" + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"data": map[string]any{
			"id":      id,
			"message": "Заказ успешно создан",
		},
	})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}
