package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	httpsrv *http.Server
	logger  *slog.Logger
}

func NewServer(httpsrv *http.Server, logger *slog.Logger) *Server {
	return &Server{
		httpsrv: httpsrv,
		logger:  logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting http server", "addr", s.httpsrv.Addr)
	if err := s.httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpsrv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	s.logger.Info("server stopped")
	return nil
}
