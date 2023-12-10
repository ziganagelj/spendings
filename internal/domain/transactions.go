package domain

import (
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type Transactions []*Transaction

func NewTransactions(transactions []Transaction) *Transactions {
	t := make(Transactions, len(transactions))
	for i, _ := range transactions {
		t[i] = &transactions[i]
	}
	return &t
}

func (t *Transactions) Add(description string, account *Account, amount float64, category Category, createdAt time.Time) *Transaction {
	transaction := NewTransaction(description, account, amount, category, createdAt)
	*t = append(*t, transaction)
	return transaction
}

func (t *Transactions) Remove(id uuid.UUID) {
	index := t.indexOf(id)
	if index == -1 {
		return
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
}

func (t *Transactions) Update(id uuid.UUID, description string, account *Account, amount float64, category Category, createdAt time.Time) *Transaction {
	index := t.indexOf(id)
	if index == -1 {
		return nil
	}
	transaction := (*t)[index]
	transaction.Update(description, account, amount, category, createdAt)

	return transaction
}

func (t *Transactions) Search(search string, month int, year int) []*Transaction {
	list := make([]*Transaction, 0)
	searchLower := strings.ToLower(search)

	for _, transaction := range *t {
		if strings.Contains(strings.ToLower(transaction.Description), searchLower) ||
			strings.Contains(strings.ToLower(string(transaction.Category)), searchLower) {
			if (month == 0 || transaction.TransactionAt.Month() == time.Month(month)) &&
				(year == 0 || transaction.TransactionAt.Year() == year) {
				list = append(list, transaction)
			}
		}
	}
	return list
}

func (t *Transactions) All() []*Transaction {
	list := make([]*Transaction, len(*t))
	copy(list, *t)
	return list
}

func (t *Transactions) Get(id uuid.UUID) *Transaction {
	index := t.indexOf(id)
	if index == -1 {
		return nil
	}
	return (*t)[index]
}

func (t *Transactions) Reorder(ids []uuid.UUID) []*Transaction {
	newTransactions := make([]*Transaction, len(ids))
	for i, id := range ids {
		newTransactions[i] = (*t)[t.indexOf(id)]
	}
	copy(*t, newTransactions)
	return newTransactions
}

func (t *Transactions) AggregateAmountsByCategory(search string, month int, year int) map[Category]float64 {
	sums := make(map[Category]float64)

	for _, t := range t.Search(search, month, year) {
		sums[t.Category] += t.Amount
	}
	return sums
}

func (t *Transactions) AggregateAmountsByCategoryAndMonth(search string, year int) []TransactionAggCategoryMonth {
	aggregation := map[string]*TransactionAggCategoryMonth{}

	for _, transaction := range t.Search(search, 0, year) {
		month := int(transaction.TransactionAt.Month())
		key := string(transaction.Category) + "-" + strconv.Itoa(month)

		if _, exists := aggregation[key]; !exists {
			aggregation[key] = &TransactionAggCategoryMonth{
				Category: transaction.Category,
				Month:    month,
				Amount:   0,
			}
		}
		aggregation[key].Amount += transaction.Amount
	}

	// Convert the map to a slice
	var aggregatedData []TransactionAggCategoryMonth
	for _, agg := range aggregation {
		aggregatedData = append(aggregatedData, *agg)
	}

	return aggregatedData
}

//func (t *Transactions) Reorder() {
//	sort.Slice(*t, func(i, j int) bool {
//		return (*t)[i].TransactionAt.Before((*t)[j].TransactionAt)
//	})
//}

func (t *Transactions) indexOf(id uuid.UUID) int {
	for i, transaction := range *t {
		if transaction.ID == id {
			return i
		}
	}
	return -1
}
