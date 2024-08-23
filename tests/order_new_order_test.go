package tests

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/tlb-katia/PET_Order_Pizza_App/tests/suite"
	pizza_orderv1 "github.com/tlb-katia/protos/protos/gen/go/pizza-order"
	"testing"
)

type pizzaType string
type pizzaSize string
type pizzaToppings []string

func (p *pizzaType) Fake() string {
	pizzaTypes := []string{"Маргарита", "Пепперони", "Гавайская", "Четыре сыра", "Капричоза"}
	return gofakeit.RandomString(pizzaTypes)
}

func (p *pizzaSize) Fake() pizza_orderv1.PizzaSize {
	pizzaSizes := []pizza_orderv1.PizzaSize{
		pizza_orderv1.PizzaSize_SMALL,
		pizza_orderv1.PizzaSize_MEDIUM,
		pizza_orderv1.PizzaSize_LARGE,
	}

	randomIndex := gofakeit.Number(0, len(pizzaSizes)-1)

	return pizzaSizes[randomIndex]
}

func (pt *pizzaToppings) Fake() []string {
	toppings := []string{
		"Extra Cheese",
		"Pepperoni",
		"Mushrooms",
		"Onions",
		"Sausage",
		"Bacon",
		"Black Olives",
		"Green Peppers",
		"Spinach",
		"Garlic",
		"Tomatoes",
		"Artichokes",
		"Chicken",
	}
	gofakeit.ShuffleStrings(toppings)
	numToppings := gofakeit.Number(1, len(toppings))

	return toppings[:numToppings]
}

func TestNewOrder(t *testing.T) {
	ctx, st := suite.New(t)

	customerName := gofakeit.Name()
	var p pizzaType
	randomPizza := p.Fake()
	var ps pizzaSize
	randomPizzaSize := ps.Fake()
	var pizzaToppings pizzaToppings
	randomToppingsSize := pizzaToppings.Fake()

	_, err := st.OrderClient.PlaceOrder(ctx, &pizza_orderv1.OrderRequest{
		CustomerName: customerName,
		PizzaType:    randomPizza,
		Size:         randomPizzaSize,
		Toppings:     randomToppingsSize,
	})
	require.NoError(t, err)
}

func TestNewOrder_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		testName     string
		customerName string
		pizzaType    pizzaType
		pizzaSize    pizza_orderv1.PizzaSize
		toppings     pizzaToppings
		expectedErr  string
	}{
		{
			testName:     "New Order With Empty Customer Name and Pizza type",
			customerName: "",
			pizzaType:    "",
			pizzaSize:    pizza_orderv1.PizzaSize_LARGE,
			toppings:     pizzaToppings{"Extra Cheese"},
			expectedErr:  "",
		},
		{
			testName:     "New Order With Empty toppings",
			customerName: "Katia",
			pizzaType:    "Neopolitana",
			pizzaSize:    pizza_orderv1.PizzaSize_SMALL,
			expectedErr:  "",
		},
		{
			testName:     "New Order With pizza size out of range",
			customerName: "Katia",
			pizzaType:    "Neopolitana",
			pizzaSize:    pizza_orderv1.PizzaSize_SMALL,
			toppings:     pizzaToppings{"Extra Cheese"},
			expectedErr:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := st.OrderClient.PlaceOrder(ctx, &pizza_orderv1.OrderRequest{
				CustomerName: tt.customerName,
				PizzaType:    string(tt.pizzaType),
				Size:         tt.pizzaSize,
				Toppings:     tt.toppings,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
