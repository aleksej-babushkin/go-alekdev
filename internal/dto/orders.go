package dto

type CreateOrderRequest struct {
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Message *string `json:"message"`
}
