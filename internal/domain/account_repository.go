package domain

import "github.com/google/uuid"

type AccountRepository interface {
	//Add(description string, account *Account, amount float64, category Category, createdAt time.Time) *Transaction
	//Remove(id uuid.UUID)
	//Update(id uuid.UUID, description string, account *Account, amount float64, category Category, createdAt time.Time) *Transaction
	Search(search string) []*Account
	//AggregateAmountsByCategory(search string) map[Category]float64
	//All() []*Account
	Get(id uuid.UUID) *Account
	//Reorder(ids []uuid.UUID) []*Transaction
}
