package postgresrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresOrderRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresOrderRepository(conn *pgxpool.Pool) *PostgresOrderRepository {
	return &PostgresOrderRepository{conn: conn}
}

func (p *PostgresOrderRepository) GetOrdersByParams(ctx context.Context, params models.QueryParams) ([]models.Order, error) {
	query := `SELECT
		id,
		user_id,
		user_email,
		user_name,
		quantity,
		subtotal,
		discount_total,
		total,
		shipping_address,
		status,
		created_at,
		updated_at
	FROM tb_orders
	LIMIT $1
	OFFSET $2`

	rows, err := p.conn.Query(ctx, query, params.Limit, params.Offset)

	orders := []models.Order{}

	if err != nil {
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var order models.Order

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.UserEmail,
			&order.UserName,
			&order.Quantity,
			&order.Subtotal,
			&order.DiscountTotal,
			&order.Total,
			&order.ShippingAddress,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)

		if err != nil {
			return orders, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
