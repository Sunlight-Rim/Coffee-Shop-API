package rest

// easyjson:json
type Coffee struct {
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Weight      uint64  `json:"weight"`
	Price       float64 `json:"price"`
}

// easyjson:json
type GetCoffeeRes struct {
	Coffee Coffee `json:"coffee"`
}

// easyjson:json
type ListCoffeeRes struct {
	CoffeeList []Coffee `json:"coffee_list"`
}
