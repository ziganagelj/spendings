package domain

import (
	"github.com/google/uuid"
	"time"
)

// The TransactionRepository interface serves as a contract or blueprint that Transactions fulfills.
// It's not directly referenced within the Transactions code itself, but it's intended to be used
// as the type for dependency injection in other parts of your application, making those parts agnostic
// to the specific implementation of transaction handling. This approach enhances modularity, testability,
// and maintainability of your code.
type TransactionRepository interface {
	Add(description string, account *Account, amount float64, category Category, createdAt time.Time) *Transaction
	Remove(id uuid.UUID)
	Update(id uuid.UUID, description string, account *Account, amount float64, category Category, createdAt time.Time) *Transaction
	Search(search string, month int, year int) []*Transaction
	AggregateAmountsByCategory(search string, month int, year int) map[Category]float64
	AggregateAmountsByCategoryAndMonth(search string, year int) []TransactionAggCategoryMonth
	All() []*Transaction
	Get(id uuid.UUID) *Transaction
	Reorder(ids []uuid.UUID) []*Transaction
}
