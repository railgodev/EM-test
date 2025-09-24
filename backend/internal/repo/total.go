package repo

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (r *Repo) Total(ctx context.Context, SubscriptionsTotal model.SubscriptionsTotal) (int, error) {
	var total int
	query := `
		SELECT SUM(price)
		FROM subscriptions
		WHERE start_date BETWEEN $1 AND $2
		AND service_name = $3
		AND user_id = $4`
	err := r.conn.QueryRow(ctx, query,
		SubscriptionsTotal.Start,
		SubscriptionsTotal.End,
		SubscriptionsTotal.ServiceName,
		SubscriptionsTotal.UserID,
	).Scan(&total)
	if err != nil {
		r.log.Info("failed to calculate total from storage", slog.String("error", err.Error()))
		return 0, errors.Wrap(err, "failed to calculate total from storage")
	}
	return total, nil
}
