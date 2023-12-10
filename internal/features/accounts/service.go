package accounts

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"spendings/internal/domain"
)

type (
	Service interface {
		// Add adds a transaction to the list
		//Add(ctx context.Context, description string, accountID uuid.UUID, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error)
		// Remove removes a transaction from the list
		//Remove(ctx context.Context, id uuid.UUID) error
		// Update updates a transaction in the list
		//Update(ctx context.Context, id uuid.UUID, description string, accountID uuid.UUID, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error)
		// Search returns a list of transactions that match the search string
		Search(ctx context.Context, search string) ([]*domain.Account, error)
		// AggregateAmountsByCategory aggregates all amounts by category
		//AggregateAmountsByCategory(ctx context.Context, search string) (map[domain.Category]float64, error)
		//// Get returns a transaction by id
		Get(ctx context.Context, id uuid.UUID) (*domain.Account, error)
		//// Sort sorts the transactions by the given ids
		//Sort(ctx context.Context, ids []uuid.UUID) error
	}

	service struct {
		db       *gorm.DB
		accounts domain.AccountRepository
	}
)

func NewService(db *gorm.DB, accounts domain.AccountRepository) Service {
	return &service{
		db:       db,
		accounts: accounts,
	}
}

func (s service) Search(_ context.Context, search string) ([]*domain.Account, error) {
	transaction := s.accounts.Search(search)
	return transaction, nil
}

//func (s service) Add(_ context.Context, description string, accountID uuid.UUID, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error) {
//	transaction := s.transactions.Add(description, accountID, amount, category, createdAt)
//	//s.db.Create(transaction)
//
//	return transaction, nil
//}
//
//func (s service) Remove(_ context.Context, id uuid.UUID) error {
//	s.transactions.Remove(id)
//	s.db.Delete(&domain.Transaction{}, id)
//
//	return nil
//}
//
//func (s service) Update(_ context.Context, id uuid.UUID, description string, accountID uuid.UUID, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error) {
//	transaction := s.transactions.Update(id, description, accountID, amount, category, createdAt)
//
//	var t domain.Transaction
//	s.db.First(&t, id)
//	s.db.Model(&t).Updates(transaction)
//
//	return transaction, nil
//}

//	func (s service) AggregateAmountsByCategory(_ context.Context, search string) (map[domain.Category]float64, error) {
//		agg := s.transactions.AggregateAmountsByCategory(search)
//		return agg, nil
//	}
func (s service) Get(_ context.Context, id uuid.UUID) (*domain.Account, error) {
	account := s.accounts.Get(id)

	return account, nil
}

//
//func (s service) Sort(_ context.Context, ids []uuid.UUID) error {
//	s.transactions.Reorder(ids)
//
//	return nil
//}
