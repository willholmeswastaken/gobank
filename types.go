package main

import "math/rand"

type Account struct {
	Id            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber int64  `json:"account_number"`
	Balance       int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:            rand.Intn(10000),
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: rand.Int63n(10000000),
	}
}
