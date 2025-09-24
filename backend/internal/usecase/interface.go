package usecase

import (
	"context"

	"github.com/railgodev/EM-test/backend/internal/model"
)

type SubscriptionsProvider interface {
	GetByID(ctx context.Context, id string) (*model.Subscription, error)
	Create(ctx context.Context, SubscriptionCreate *model.SubscriptionCreate) (*model.Subscription, error)
	List(ctx context.Context, page, pageSize int) ([]*model.Subscription, error)
	Update(ctx context.Context, id string, SubscriptionUpdate *model.SubscriptionUpdate) (*model.Subscription, error)
	Delete(ctx context.Context, id string) (*model.Subscription, error)
	Total(ctx context.Context, SubscriptionsTotal model.SubscriptionsTotal) (int, error)
}
