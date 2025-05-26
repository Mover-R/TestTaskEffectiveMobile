package main

import (
	"TestTaskEffectiveMobile/internal/config"
	server "TestTaskEffectiveMobile/internal/transport/rest"
	"TestTaskEffectiveMobile/pkg/logger"
	"TestTaskEffectiveMobile/pkg/postgres"
	"context"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	ctx, _ = logger.NewLogger(ctx)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	config, err := config.NewConfig(ctx)

	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Failed load logger", zap.Error(err))
	}
	logger.GetLoggerFromCtx(ctx).Info(ctx, "Succesfully load config")

	pgDB, err := postgres.NewPostgres(ctx, &config.PostgresCFG)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Failsed connect to postgres DB", zap.Error(err))
	}
	if err := pgDB.Ping(ctx); err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Failed ping pgDB", zap.Error(err))
	}
	logger.GetLoggerFromCtx(ctx).Info(ctx, "Succesfully connected to pgDB")

	r, err := server.NewRouter(ctx, config)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "")
	}
	r.Run(ctx)

	<-ctx.Done()
	pgDB.Close()
	logger.GetLoggerFromCtx(ctx).Info(ctx, "Server Stopped")
}
