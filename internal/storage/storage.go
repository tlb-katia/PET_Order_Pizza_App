package storage

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
)

var (
	NewOrderPlaced     = "Nice choice! We will take care of your pizza"
	DeleteOrderMessage = "Your order is successfully deleted.\n Hope to see you again"
)
