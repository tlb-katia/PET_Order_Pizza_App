package suite

import (
	"context"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/config"
	pizza_orderv1 "github.com/tlb-katia/protos/protos/gen/go/pizza-order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
	"time"
)

type Suite struct {
	*testing.T
	Cfg         *config.Config
	OrderClient pizza_orderv1.POrderClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper() // to understand that it is a helper function - not a final one
	t.Parallel()

	conf := config.MustLoad()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	conn, err := grpc.DialContext(
		context.Background(),
		gRPCAddress(conf),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	return ctx, &Suite{
		T:           t,
		Cfg:         conf,
		OrderClient: pizza_orderv1.NewPOrderClient(conn),
	}
}

func gRPCAddress(cfg *config.Config) string {
	return net.JoinHostPort("localhost", cfg.Grpc.Port)
}
