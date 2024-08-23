package entities

type PizzaSize int32

const (
	SMALL PizzaSize = iota
	MEDIUM
	LARGE
)

type OrderStatus int32

const (
	PREPARING OrderStatus = iota
	ON_THE_WAY
	DELIVERED
	CANCELLED
)

type PizzaOrderReq struct {
	CustomerName string
	PizzaType    string
	PizzaSize    PizzaSize
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
	OrderStatus OrderStatus
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
	PizzaSize    PizzaSize
	Toppings     []string
	OrderStatus  OrderStatus
}
