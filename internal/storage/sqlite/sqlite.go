package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/entities"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	const op = "sqlite.New"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s Storage) GetOrder(ctx context.Context, pizzaID string) (*entities.PizzaDbParams, error) {
	const op = "sqlite.GetOrder"
	var order entities.PizzaDbParams

	query, err := s.db.PrepareContext(ctx, "SELECT * FROM orders WHERE pizza_id=?")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	res := query.QueryRowContext(ctx, pizzaID)
	err = res.Scan(order.OrderId, order.CustomerName, order.PizzaType, order.PizzaSize, order.OrderStatus)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &order, nil
}

func (s Storage) SaveOrder(ctx context.Context, req *entities.PizzaOrderReq) (string, error) {
	const op = "sqlite.SaveOrder"
	var lastID string

	queryOrder, err := s.db.PrepareContext(ctx,
		"INSERT INTO orders (customer_name, pizza_type, pizza_size) VALUES (?,?,?) RETURNING id")

	if err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	err = queryOrder.QueryRow(req.CustomerName, req.PizzaType, req.PizzaSize).Scan(&lastID)
	if err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	queryToppings, err := s.db.PrepareContext(ctx,
		"INSERT INTO toppings (order_id, topping) VALUES (?,?)")

	for topping := range req.Toppings {
		_, err = queryToppings.Exec(lastID, topping)
		if err != nil {
			return "", fmt.Errorf("%s, %w", op, err)
		}
	}

	return lastID, nil
}

func (s Storage) DeleteOrder(ctx context.Context, pizzaID string) error {
	const op = "sqlite.DeleteOrder"

	query, err := s.db.PrepareContext(ctx, "DELETE FROM orders WHERE id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = query.ExecContext(ctx, pizzaID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
