package models

import "errors"

type Customer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Branch  string `json:"branch"`
	Balance int    `json:"balance"`
	City    string `json:"city"`
}

type Response struct {
	Status  string     `json:"status"`
	Data    []Customer `json:"data"`
	Message string     `json:"message"`
}

var customers = []Customer{
	{Id: "first@1", Name: "first", Branch: "bangfir", Balance: 20000, City: "bengaluru"},
	{Id: "second@2", Name: "second", Branch: "bangsec", Balance: 30000, City: "bengaluru"},
	{Id: "third@3", Name: "third", Branch: "bangthird", Balance: 50000, City: "bengaluru"},
}

func Askdata() ([]Customer, error) {
	if customers != nil {
		return customers, nil
	} else {
		return nil, errors.New("data not found")
	}
}

func Setdata(out []Customer) {
	customers = out
}
