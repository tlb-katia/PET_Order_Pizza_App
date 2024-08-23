package grpc

import (
	"fmt"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/grpc/order"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       string
}

func New(log *slog.Logger, port string, orderService order.Order) *App {
	grpcServer := grpc.NewServer()
	order.Register(grpcServer, orderService)

	return &App{
		log:        log,
		grpcServer: grpcServer,
		port:       port,
	}
}

func (a *App) Run() error {
	const method = "app.Run"
	log := a.log.With(
		slog.String("method", method),
		slog.String("port", a.port))

	lis, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		log.Error("failed to listen", "err", err)
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Info("grpc server is listening", "port", a.port)

	if err := a.grpcServer.Serve(lis); err != nil {
		log.Error("failed to serve", "err", err)
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func (a *App) Stop() {
	const method = "app.Stop"
	a.log.With(
		slog.String("method", method),
		slog.String("port", a.port))

	a.grpcServer.GracefulStop()
}
