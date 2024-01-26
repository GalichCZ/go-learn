package main

import (
	"github.com/google/uuid"
)

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    string `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    uuid.NewString(),
	}
}
