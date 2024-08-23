package main

import (
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/app"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
	EnvDev   = "dev"
)

func main() {
	cfg := config.MustLoad()
	log := setUpLogger(cfg.Env)
	log.Info("Starting the server", slog.String("port", cfg.Grpc.Port))

	application := app.New(log, cfg.Grpc.Port, cfg.PathDb)

	go application.GRPCSrv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.GRPCSrv.Stop()

	log.Info("shutting down")
}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
