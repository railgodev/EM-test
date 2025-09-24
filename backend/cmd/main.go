package main

import (
	"log/slog"
	"os"

	"github.com/railgodev/EM-test/backend/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		slog.Error("app run", slog.Any("err", err))
		os.Exit(1)
	}
}
