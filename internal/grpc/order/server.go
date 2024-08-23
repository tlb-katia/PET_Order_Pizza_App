package order

import (
	"context"
	"fmt"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/entities"
	pizza_orderv1 "github.com/tlb-katia/protos/protos/gen/go/pizza-order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order interface {
	PlaceOrder(ctx context.Context, req *entities.PizzaOrderReq) (*entities.PizzaOrderResp, error)
	CheckOrderStatus(ctx context.Context, req *entities.OrderStatusRequest) (*entities.OrderStatusResp, error)
	CancelOrder(ctx context.Context, req *entities.CancelOrderRequest) (*entities.CancelOrderResponse, error)
}

type serverAPI struct {
	pizza_orderv1.UnimplementedPOrderServer
	order Order
}

func Register(grpcServer *grpc.Server, order Order) {
	pizza_orderv1.RegisterPOrderServer(grpcServer, &serverAPI{order: order})
}

func (s *serverAPI) PlaceOrder(ctx context.Context, req *pizza_orderv1.OrderRequest) (*pizza_orderv1.OrderResponse, error) {
	err := validatePizzaParams(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reqEntity := &entities.PizzaOrderReq{
		CustomerName: req.CustomerName,
		PizzaType:    req.PizzaType,
		PizzaSize:    entities.PizzaSize(req.Size),
		Toppings:     req.Toppings,
	}

	order, err := s.order.PlaceOrder(ctx, reqEntity)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	return &pizza_orderv1.OrderResponse{
			OrderId: order.OrderId,
			Message: order.Message},
		nil
}

func (s *serverAPI) CheckOrderStatus(ctx context.Context, req *pizza_orderv1.OrderStatusRequest) (*pizza_orderv1.OrderStatusResponse, error) {
	if req.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "Order ID is required")
	}

	orderStatus, err := s.order.CheckOrderStatus(ctx, &entities.OrderStatusRequest{OrderId: req.OrderId})
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	return &pizza_orderv1.OrderStatusResponse{
		OrderId: orderStatus.OrderId,
		Status:  pizza_orderv1.OrderStatus(orderStatus.OrderStatus),
	}, nil
}

func (s *serverAPI) CancelOrder(ctx context.Context, req *pizza_orderv1.CancelOrderRequest) (*pizza_orderv1.CancelOrderResponse, error) {
	if req.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "Order ID is required")
	}
	order, err := s.order.CancelOrder(ctx, &entities.CancelOrderRequest{OrderId: req.OrderId})
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	return &pizza_orderv1.CancelOrderResponse{
		OrderId: order.OrderId,
		Message: order.Message,
	}, nil
}

func validatePizzaParams(req *pizza_orderv1.OrderRequest) error {
	if req.CustomerName == "" {
		return fmt.Errorf("customer name is required")
	}
	if req.PizzaType == "" {
		return fmt.Errorf("pizza type is required")
	}
	if req.Size == 0 {
		return fmt.Errorf("pizza size is required")
	}
	if req.Toppings == nil {
		return fmt.Errorf("toppings are required")
	}
	return nil
}
