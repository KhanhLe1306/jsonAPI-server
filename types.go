package main

type Account struct {
	Id			uint 	`json:"id"`
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	Balance 	float64	`json:"balance"`
}

var Id uint = 1

func NewAccount(firstName, lastName string) *Account{
	return &Account{
		Id: Id,
		FirstName: firstName,
		LastName: lastName,
	}
}






