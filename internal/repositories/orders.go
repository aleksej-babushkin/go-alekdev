package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	ID        int64
	Name      string
	Phone     string
	Message   *string
	IP        string
	UserAgent string
	SourceURL string
}

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(
	ctx context.Context,
	order *Order,
) (int64, error) {
	var id int64

	err := r.db.QueryRow(ctx, `
		INSERT INTO lead (
			name,
			phone,
			message,
			created_at,
			ip,
			user_agent,
			source_url
		)
		VALUES ($1, $2, $3, NOW(), $4, $5, $6)
		RETURNING id
	`,
		order.Name,
		order.Phone,
		order.Message,
		order.IP,
		order.UserAgent,
		order.SourceURL,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
