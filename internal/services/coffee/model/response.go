package model

// easyjson:json
type Coffee struct {
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Weight      uint64  `json:"weight"`
	Price       float64 `json:"price"`
}

// ListCoffees

// easyjson:json
type ListCoffeesResDelivery struct {
	CoffeeList []Coffee `json:"coffee_list"`
}

type ListCoffeesResUsecase struct {
	CoffeeList []Coffee
}

type ListCoffeesResStorage struct {
	CoffeeList []Coffee
}

// GetCoffee

// easyjson:json
type GetCoffeeInfoResDelivery struct {
	Coffee *Coffee `json:"coffee"`
}

type GetCoffeeInfoResUsecase struct {
	Coffee *Coffee
}

type GetCoffeeInfoResStorage struct {
	Coffee *Coffee
}

// ListCoffees

// easyjson:json
type ListToppingsResDelivery struct {
	ToppingsList []string `json:"toppings_list"`
}

type ListToppingsResUsecase struct {
	ToppingsList []string
}

type ListToppingsResStorage struct {
	ToppingsList []string
}
