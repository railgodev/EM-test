package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/EM-test/backend/config"
	"github.com/railgodev/EM-test/backend/internal/handler"
	"github.com/railgodev/EM-test/backend/internal/logger"
	"github.com/railgodev/EM-test/backend/internal/repo"
	"github.com/railgodev/EM-test/backend/internal/server"
	"github.com/railgodev/EM-test/backend/internal/storage"
	"github.com/railgodev/EM-test/backend/internal/usecase"
	"github.com/railgodev/EM-test/backend/migrate"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.LoadConfig()

	logger := logger.Setup(cfg.App.LogLevel)

	if err := migrate.Run(cfg.GetDSN(), cfg.MigratePath, logger); err != nil {
		logger.Error("migrate", slog.Any("err", err))
		return err
	}

	conn, err := storage.GetConn(logger, cfg.GetConnStr())
	if err != nil {
		logger.Error("connection pool", slog.Any("err", err))
		return err
	}
	repo := repo.New(conn, logger)

	uc := usecase.New(repo)

	h := handler.New(uc, logger)

	if cfg.App.LogLevel != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Serve swagger.yaml
	router.StaticFile("/swagger.yaml", "./swagger.yaml")

	// Serve Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	api := router.Group("/api/v1/")
	api.GET("/subscriptions", h.List)
	api.POST("/subscriptions", h.Create)
	api.GET("/subscriptions/:id", h.GetByID)
	api.PUT("/subscriptions/:id", h.Update)
	api.DELETE("/subscriptions/:id", h.Delete)
	api.GET("/subscriptions/total", h.Total)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.App.Address, cfg.App.Port),
		Handler: router,
	}

	server := server.NewServer(srv, logger)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Start(ctx); err != nil {
			logger.Error("failed to start server", slog.Any("err", err))
		}
	}()

	<-done

	if err := server.Stop(); err != nil {
		logger.Error("failed to stop server", slog.Any("err", err))
		return err
	}
	return nil
}
