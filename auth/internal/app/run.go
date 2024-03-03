package app

import (
	"context"
	"fmt"

	"auth/internal/config"
	"auth/internal/handler"
	"auth/internal/log"
	"auth/internal/migration"
	"auth/internal/service"
	"auth/internal/store"
)

func Run(appName string) error {
	ctx := context.Background()

	_, err := initPaas(ctx, appName, "configs/web")
	if err != nil {
		return err
	}

	if err := migration.Up(ctx, config.Default().AuthDB.MigrationURL); err != nil {
		return err
	}

	dbPool, err := store.NewDBPool(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDBPool, err.Error())
	}

	authStore := store.New(dbPool)
	authService := service.NewAuthService(authStore)
	srv := handler.NewAuthGRPCHandler(authService)

	grpcServer, grpcListener, err := handler.RegisterGRPCServer(srv)
	if err != nil {
		return err
	}

	go func() {
		log.Default().Info().Msgf("starting grpc server on port %v", config.Default().GRPC.Port)

		if err = grpcServer.Serve(grpcListener); err != nil {
			log.Default().Err(err).Msg("grpc error")
		}
	}()

	<-shutdown(ctx)

	return nil
}
