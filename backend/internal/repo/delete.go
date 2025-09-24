package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/railgodev/EM-test/backend/internal/model"
)

func (r *Repo) Delete(ctx context.Context, id string) (*model.Subscription, error) {
	subscription, err := r.GetByID(ctx, id)
	if subscription == nil && err == nil {
		return nil, nil
	}

	query := `DELETE 
			FROM subscriptions 
			WHERE id = $1`
	_, err = r.conn.Exec(ctx, query, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete subscription from storage")
	}
	return subscription, nil
}
