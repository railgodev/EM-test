package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (r *Repo) List(ctx context.Context, page, pageSize int) ([]*model.Subscription, error) {
	var subscriptions []*model.Subscription

	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions
			  ORDER BY start_date
			  LIMIT $1 OFFSET $2`
	rows, err := r.conn.Query(ctx, query, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list subscriptions from storage")
	}
	defer rows.Close()

	for rows.Next() {
		var subscription model.Subscription
		if err := rows.Scan(
			&subscription.ID,
			&subscription.ServiceName,
			&subscription.Price,
			&subscription.UserID,
			&subscription.StartDate,
			&subscription.EndDate,
		); err != nil {
			return nil, errors.Wrap(err, "failed to scan subscription row")
		}
		subscriptions = append(subscriptions, &subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to iterate subscription rows")
	}

	return subscriptions, nil
}
