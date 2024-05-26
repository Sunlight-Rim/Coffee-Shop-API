package model

// GetCoffee

type GetCoffeeRes struct {
	Coffee *Coffee
}

type StorageGetCoffeeRes struct {
	Coffee *Coffee
}

// ListCoffee

type ListCoffeeRes struct {
	CoffeeList []Coffee
}

type StorageListCoffeeRes struct {
	CoffeeList []Coffee
}
