package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (r *Repo) Create(ctx context.Context, SubscriptionCreate *model.SubscriptionCreate) (*model.Subscription, error) {
	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, service_name, price, user_id, start_date, end_date`
	subscription := &model.Subscription{}
	err := r.conn.QueryRow(ctx, query,
		SubscriptionCreate.ServiceName,
		SubscriptionCreate.Price,
		SubscriptionCreate.UserID,
		SubscriptionCreate.StartDate,
		SubscriptionCreate.EndDate,
	).Scan(
		&subscription.ID,
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserID,
		&subscription.StartDate,
		&subscription.EndDate,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subscription in storage")
	}

	return subscription, nil
}
