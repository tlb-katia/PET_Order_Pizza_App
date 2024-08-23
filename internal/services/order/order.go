package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/entities"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/storage"
	"log/slog"
)

type Order struct {
	log      *slog.Logger
	provider Provider
}

type Provider interface {
	GetOrder(ctx context.Context, pizzaID string) (*entities.PizzaDbParams, error)
	SaveOrder(ctx context.Context, req *entities.PizzaOrderReq) (string, error) //id
	DeleteOrder(ctx context.Context, pizzaID string) error
}

func New(log *slog.Logger, provider Provider) *Order {
	return &Order{log, provider}
}

func (o *Order) PlaceOrder(ctx context.Context, req *entities.PizzaOrderReq) (*entities.PizzaOrderResp, error) {
	const op = "Order.PlaceOrder"
	log := o.log.With(slog.String("operation", op),
		slog.String("customerName", req.CustomerName))
	log.Info("placing a new order")

	orderId, err := o.provider.SaveOrder(ctx, req)
	if err != nil {
		log.Error("failed to place an order", err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &entities.PizzaOrderResp{
		OrderId: orderId,
		Message: storage.NewOrderPlaced,
	}, nil
}

func (o *Order) CheckOrderStatus(ctx context.Context, req *entities.OrderStatusRequest) (*entities.OrderStatusResp, error) {
	const op = "Order.CheckOrderStatus"
	log := o.log.With(slog.String("operation", op),
		slog.String("orderId", req.OrderId))
	log.Info("checking order status")

	order, err := o.provider.GetOrder(ctx, req.OrderId)
	if err != nil {
		log.Error("failed to find order", errors.New("invalid order id"))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &entities.OrderStatusResp{
		OrderId:     order.OrderId,
		OrderStatus: order.OrderStatus,
	}, nil
}

func (o *Order) CancelOrder(ctx context.Context, req *entities.CancelOrderRequest) (*entities.CancelOrderResponse, error) {
	const op = "Order.CancelOrder"
	log := o.log.With(slog.String("operation", op),
		slog.String("orderId", req.OrderId))
	log.Info("canceling order")

	err := o.provider.DeleteOrder(ctx, req.OrderId)
	if err != nil {
		log.Error("failed to cancel order", errors.New("check credentials"))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &entities.CancelOrderResponse{
		OrderId: req.OrderId,
		Message: storage.DeleteOrderMessage,
	}, nil
}
