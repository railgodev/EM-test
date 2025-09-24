package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (r *Repo) GetByID(ctx context.Context, id string) (*model.Subscription, error) {
	var subscription model.Subscription

	query := `SELECT id, service_name, price, user_id, start_date, end_date 
			FROM subscriptions 
			WHERE id = $1`

	err := r.conn.QueryRow(ctx, query, id).Scan(
		&subscription.ID,
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserID,
		&subscription.StartDate,
		&subscription.EndDate,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get subscription from storage")
	}

	return &subscription, nil
}
