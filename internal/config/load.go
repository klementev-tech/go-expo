package config

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flag      = "conf"
	envPrefix = "GE"
)

func Load() (*Config, error) {
	pflag.String(flag, "", "config file path")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, err
	}

	bindEnv("log.level", "LOG_LEVEL")

	viper.SetConfigFile(viper.GetString(flag))

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("config file not found: %w", err)
		}

		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return cfg, nil
}

func bindEnv(key, env string) {
	err := viper.BindEnv(key, fmt.Sprintf("%s_%s", envPrefix, env))
	if err != nil {
		slog.Default().WarnContext(context.Background(), "error binding env", "error", err)
	}
}
