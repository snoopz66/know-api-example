package models

type User struct {
	UUID    string   `json:"uuid"`
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Email   string   `json:"email"`
	Address *Address `json:"address"`
}

type Address struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}
