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

// ListCoffee

// easyjson:json
type ListCoffeeResDelivery struct {
	CoffeeList []Coffee `json:"coffee_list"`
}

type ListCoffeeResUsecase struct {
	CoffeeList []Coffee
}

type ListCoffeeResStorage struct {
	CoffeeList []Coffee
}
