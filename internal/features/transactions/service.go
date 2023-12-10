package transactions

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"spendings/internal/domain"
	"time"
)

type (
	Service interface {
		// Add adds a transaction to the list
		Add(ctx context.Context, description string, account *domain.Account, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error)
		// Remove removes a transaction from the list
		Remove(ctx context.Context, id uuid.UUID) error
		// Update updates a transaction in the list
		Update(ctx context.Context, id uuid.UUID, description string, account *domain.Account, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error)
		// Search returns a list of transactions that match the search string
		Search(ctx context.Context, search string, month int, year int) ([]*domain.Transaction, error)
		// AggregateAmountsByCategory aggregates all amounts by category
		AggregateAmountsByCategory(ctx context.Context, search string, month int, year int) (map[domain.Category]float64, error)
		AggregateAmountsByCategoryAndMonth(ctx context.Context, search string, year int) ([]domain.TransactionAggCategoryMonth, error)
		// Get returns a transaction by id
		Get(ctx context.Context, id uuid.UUID) (*domain.Transaction, error)
		// Sort sorts the transactions by the given ids
		Sort(ctx context.Context, ids []uuid.UUID) error
	}

	service struct {
		db           *gorm.DB
		transactions domain.TransactionRepository
	}
)

func NewService(db *gorm.DB, transactions domain.TransactionRepository) Service {
	return &service{
		db:           db,
		transactions: transactions,
	}
}

func (s service) Add(_ context.Context, description string, account *domain.Account, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error) {
	transaction := s.transactions.Add(description, account, amount, category, createdAt)
	s.db.Create(transaction)

	return transaction, nil
}

func (s service) Remove(_ context.Context, id uuid.UUID) error {
	s.transactions.Remove(id)
	s.db.Delete(&domain.Transaction{}, id)

	return nil
}

func (s service) Update(_ context.Context, id uuid.UUID, description string, account *domain.Account, amount float64, category domain.Category, createdAt time.Time) (*domain.Transaction, error) {
	transaction := s.transactions.Update(id, description, account, amount, category, createdAt)

	var t domain.Transaction
	s.db.First(&t, id)
	s.db.Model(&t).Updates(transaction)

	return transaction, nil
}

func (s service) Search(_ context.Context, search string, month int, year int) ([]*domain.Transaction, error) {
	transaction := s.transactions.Search(search, month, year)
	return transaction, nil
}

func (s service) AggregateAmountsByCategory(_ context.Context, search string, month int, year int) (map[domain.Category]float64, error) {
	agg := s.transactions.AggregateAmountsByCategory(search, month, year)
	return agg, nil
}

func (s service) AggregateAmountsByCategoryAndMonth(_ context.Context, search string, year int) ([]domain.TransactionAggCategoryMonth, error) {
	agg := s.transactions.AggregateAmountsByCategoryAndMonth(search, year)
	return agg, nil
}

func (s service) Get(_ context.Context, id uuid.UUID) (*domain.Transaction, error) {
	transaction := s.transactions.Get(id)

	return transaction, nil
}

func (s service) Sort(_ context.Context, ids []uuid.UUID) error {
	s.transactions.Reorder(ids)

	return nil
}
