package app

import (
	"context"
	"errors"

	"auth/internal/config"
	"auth/internal/log"
)

type platform struct {
}

var (
	ErrLogLevelParse = errors.New("log level parse error")
)

func initPaas(ctx context.Context, appName, metricsDir string) (*platform, error) {
	if err := config.LoadConfig(); err != nil {
		return nil, err
	}

	if err := log.SetupLogger(appName); err != nil {
		return nil, err
	}

	log.Default().Info().Msg("logger and configuration: ok")

	p := new(platform)

	return p, nil

}
