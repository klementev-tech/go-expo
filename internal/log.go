package internal

import (
	"log/slog"
	"os"

	"github.com/klementev-tech/go-expo/internal/config"
)

func InitLog(cfg config.Log) error {
	var l slog.Level

	if err := l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return err
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: l,
	})

	slog.SetDefault(slog.New(h))
	return nil
}
