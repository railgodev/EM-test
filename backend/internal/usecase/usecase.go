package usecase

import (
	"context"

	"github.com/railgodev/EM-test/backend/internal/model"
	"github.com/railgodev/EM-test/backend/internal/repo"
)

type Usecase struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetByID(ctx context.Context, id string) (*model.Subscription, error) {
	return u.repo.GetByID(ctx, id)
}
func (u *Usecase) Create(ctx context.Context, SubscriptionCreate *model.SubscriptionCreate) (*model.Subscription, error) {
	return u.repo.Create(ctx, SubscriptionCreate)
}
func (u *Usecase) List(ctx context.Context, page, pageSize int) ([]*model.Subscription, error) {
	return u.repo.List(ctx, page, pageSize)
}
func (u *Usecase) Update(ctx context.Context, id string, SubscriptionUpdate *model.SubscriptionUpdate) (*model.Subscription, error) {
	return u.repo.Update(ctx, id, SubscriptionUpdate)
}
func (u *Usecase) Delete(ctx context.Context, id string) (*model.Subscription, error) {
	return u.repo.Delete(ctx, id)
}
func (u *Usecase) Total(ctx context.Context, SubscriptionsTotal model.SubscriptionsTotal) (int, error) {
	return u.repo.Total(ctx, SubscriptionsTotal)
}
