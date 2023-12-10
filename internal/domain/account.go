package domain

import (
	"github.com/google/uuid"
	"spendings/internal/db"
	"spendings/internal/util"
)

type Account struct {
	db.Model
	Name     string
	Currency currency
	Amount   float64
}

func NewAccount(name string, currency currency, amount float64) *Account {
	return &Account{
		Model: db.Model{
			ID: uuid.New(),
		},
		Name:     name,
		Currency: currency,
		Amount:   amount,
	}
}

func (t *Account) Update(name string, currency currency, amount float64) {
	t.Name = name
	t.Currency = currency
	t.Amount = amount
}

func (t *Account) AmountString() string {
	return util.FormatFloat(t.Amount)
}
