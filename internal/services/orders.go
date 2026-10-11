package services

import (
	"context"
	"errors"
	"strings"

	"app/internal/repositories"
)

var ErrInvalidOrder = errors.New("name and phone are required")

type CreateOrderInput struct {
	Name      string
	Phone     string
	Message   *string
	IP        string
	UserAgent string
	SourceURL string
}

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(
	ctx context.Context,
	input CreateOrderInput,
) (int64, error) {
	input.Name = strings.TrimSpace(input.Name)
	input.Phone = strings.TrimSpace(input.Phone)

	if input.Name == "" || input.Phone == "" {
		return 0, ErrInvalidOrder
	}

	if len(input.Name) > 255 ||
		len(input.Phone) > 255 ||
		(input.Message != nil && len(*input.Message) > 1000) ||
		len(input.UserAgent) > 255 ||
		len(input.SourceURL) > 500 {
		return 0, ErrInvalidOrder
	}

	order := &repositories.Order{
		Name:      input.Name,
		Phone:     input.Phone,
		Message:   input.Message,
		IP:        input.IP,
		UserAgent: input.UserAgent,
		SourceURL: input.SourceURL,
	}

	return s.repo.Create(ctx, order)
}
