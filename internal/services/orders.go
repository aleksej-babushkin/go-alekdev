package services

type Order struct {
	ID int64 `json:"id"`
}

type OrderService struct {
	// Repository
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) Create() (*Order, error) {
	order := &Order{
		ID: 1,
	}

	return order, nil
}
