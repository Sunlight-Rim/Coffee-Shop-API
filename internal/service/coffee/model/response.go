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
type DeliveryGetCoffeeInfoRes struct {
	Coffee *Coffee `json:"coffee"`
}

type UsecaseGetCoffeeInfoRes struct {
	Coffee *Coffee
}

type StorageGetCoffeeInfoRes struct {
	Coffee *Coffee
}

// ListCoffee

// easyjson:json
type DeliveryListCoffeeRes struct {
	CoffeeList []Coffee `json:"coffee_list"`
}

type UsecaseListCoffeeRes struct {
	CoffeeList []Coffee
}

type StorageListCoffeeRes struct {
	CoffeeList []Coffee
}
