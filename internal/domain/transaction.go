package domain

import (
	"github.com/google/uuid"
	"spendings/internal/db"
	"spendings/internal/util"
	"time"
)

type Transaction struct {
	db.Model
	Description   string
	AccountID     uuid.UUID
	Account       *Account
	Amount        float64
	Category      Category
	TransactionAt time.Time
}

func NewTransaction(description string, account *Account, amount float64, category Category, transactionAt time.Time) *Transaction {
	return &Transaction{
		Model: db.Model{
			ID: uuid.New(),
		},
		Description:   description,
		AccountID:     account.ID,
		Account:       account,
		Amount:        amount,
		Category:      category,
		TransactionAt: transactionAt,
	}
}

func (t *Transaction) Update(description string, account *Account, amount float64, category Category, transactionAt time.Time) {
	t.Description = description
	t.AccountID = account.ID
	t.Account = account
	t.Amount = amount
	t.Category = category
	t.TransactionAt = transactionAt
}

func (t *Transaction) AmountString() string {
	return util.FormatFloat(t.Amount)
}

func (t *Transaction) TransactionAtString() string {
	return t.TransactionAt.Format(util.DateFormat)
}

type TransactionAggCategoryMonth struct {
	Category Category
	Month    int
	Amount   float64
}
