package entities

import pizza_orderv1 "github.com/tlb-katia/protos/protos/gen/go/pizza-order"

type PizzaOrderReq struct {
	CustomerName string
	PizzaType    string
	PizzaSize    pizza_orderv1.PizzaSize
	Toppings     []string
}

type PizzaOrderResp struct {
	OrderId string
	Message string
}

type OrderStatusRequest struct {
	OrderId string
}

type OrderStatusResp struct {
	OrderId     string
	OrderStatus pizza_orderv1.OrderStatus
}

type CancelOrderRequest struct {
	OrderId string
}

type CancelOrderResponse struct {
	OrderId string
	Message string
}

type PizzaDbParams struct {
	OrderId      string
	CustomerName string
	PizzaType    string
	PizzaSize    pizza_orderv1.PizzaSize
	Toppings     []string
	OrderStatus  pizza_orderv1.OrderStatus
}
