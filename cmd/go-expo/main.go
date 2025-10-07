package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/klementev-tech/go-expo/internal"
	"github.com/klementev-tech/go-expo/internal/config"
)

func run() int {
	cfg, err := config.Load()
	if err != nil {
		slog.Default().ErrorContext(context.Background(), "failed to load config", "error", err)
		return 1
	}

	err = internal.InitLog(cfg.Log)
	if err != nil {
		slog.Default().ErrorContext(context.Background(), "failed to init logger", "error", err)
		return 1
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	slog.Default().InfoContext(ctx, "starting go-expo application")

	slog.Default().DebugContext(ctx, "config param", slog.String("log_level", cfg.Log.Level))

	<-ctx.Done()

	slog.Default().InfoContext(ctx, "go-expo completed", "error", ctx.Err())
	return 0
}

func main() {
	os.Exit(run())
}
