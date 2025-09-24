package handler

import (
	"log/slog"

	"github.com/railgodev/EM-test/backend/internal/usecase"
)

type Handle struct {
	uc  usecase.SubscriptionsProvider
	log *slog.Logger
}

func New(uc usecase.SubscriptionsProvider, log *slog.Logger) *Handle {
	return &Handle{
		uc:  uc,
		log: log,
	}
}
