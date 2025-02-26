package main

import (
	"context"
	"country_service/internal/app"
	"country_service/internal/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// TODO start microservice
func main() {
	//TODO init config obj
	cfg := config.MustLoad()

	//TODO init logger
	//if need more opt make create in lib/logger
	log := setupLogger()

	//TODO init app
	ctx := context.Background()
	app := app.New(ctx, log, cfg.GRPC.Port, cfg.Dsn, cfg.TokenTTL)
	log.Info("Start service")
	//TODO start grpc-Server

	app.GRPCServer.MustRun()

	//TODO graceful shutdown
	// Graceful shutdown
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	// Waiting for SIGINT (pkill -2) or SIGTERM
	<-stop
	// initiate graceful shutdown
	app.GRPCServer.Stop()
	log.Info("GRPCserver stopped")
	app.Storage.Stop()
	log.Info("Postgres connection closed")
}

func setupLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}
