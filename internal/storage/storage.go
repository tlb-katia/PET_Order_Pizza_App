package storage

import "errors"

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrEmptyCustomerName  = errors.New("empty customer name")
	ErrEmptyPizzaType     = errors.New("empty pizza type")
	ErrEmptyToppings      = errors.New("empty toppings")
	ErrSizeOutOfRange     = errors.New("pizza size does not exist")
)

var (
	NewOrderPlaced     = "Nice choice! We will take care of your pizza"
	DeleteOrderMessage = "Your order is successfully deleted.\n Hope to see you again"
)
