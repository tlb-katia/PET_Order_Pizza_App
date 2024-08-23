package app

import (
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/app/grpc"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/services/order"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/storage/sqlite"
	"log/slog"
)

type App struct {
	GRPCSrv *grpc.App
}

func New(log *slog.Logger, port string, dbPath string) *App {
	storage, err := sqlite.New(dbPath)
	if err != nil {
		log.Error("Failed to init database", "error", err)
		return nil
	}

	service := order.New(log, storage)

	grpcApp := grpc.New(log, port, service)

	return &App{
		GRPCSrv: grpcApp,
	}
}
