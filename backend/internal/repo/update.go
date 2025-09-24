package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (r *Repo) Update(ctx context.Context, id string, SubscriptionUpdate *model.SubscriptionUpdate) (*model.Subscription, error) {
	setClauses := []string{}
	args := []interface{}{}
	argIdx := 1

	if SubscriptionUpdate.ServiceName != nil {
		setClauses = append(setClauses, fmt.Sprintf("service_name = $%d", argIdx))
		args = append(args, *SubscriptionUpdate.ServiceName)
		argIdx++
	}
	if SubscriptionUpdate.Price != nil {
		setClauses = append(setClauses, fmt.Sprintf("price = $%d", argIdx))
		args = append(args, *SubscriptionUpdate.Price)
		argIdx++
	}
	if SubscriptionUpdate.UserID != nil {
		setClauses = append(setClauses, fmt.Sprintf("user_id = $%d", argIdx))
		args = append(args, *SubscriptionUpdate.UserID)
		argIdx++
	}
	if SubscriptionUpdate.StartDate != nil {
		setClauses = append(setClauses, fmt.Sprintf("start_date = $%d", argIdx))
		args = append(args, *SubscriptionUpdate.StartDate)
		argIdx++
	}
	if SubscriptionUpdate.EndDate != nil {
		setClauses = append(setClauses, fmt.Sprintf("end_date = $%d", argIdx))
		args = append(args, *SubscriptionUpdate.EndDate)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, errors.New("no fields to update")
	}

	query := fmt.Sprintf(`
		UPDATE subscriptions
		SET %s
		WHERE id = $%d
		RETURNING id, service_name, price, user_id, start_date, end_date
	`, strings.Join(setClauses, ", "), argIdx)

	args = append(args, id)

	subscription := &model.Subscription{}
	err := r.conn.QueryRow(ctx, query, args...).Scan(
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
		return nil, errors.Wrap(err, "failed to update subscription in storage")
	}

	return subscription, nil
}
