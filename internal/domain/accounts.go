package domain

import (
	"github.com/google/uuid"
	"strings"
)

type Accounts []*Account

func NewAccounts(accounts []Account) *Accounts {
	a := make(Accounts, len(accounts))
	for i, _ := range accounts {
		a[i] = &accounts[i]
	}
	return &a
}

//	func (t *Transactions) Add(description string, accountID uuid.UUID, amount float64, category Category, createdAt time.Time) *Transaction {
//		transaction := NewTransaction(description, accountID, amount, category, createdAt)
//		*t = append(*t, transaction)
//		return transaction
//	}
//
//	func (t *Transactions) Remove(id uuid.UUID) {
//		index := t.indexOf(id)
//		if index == -1 {
//			return
//		}
//		*t = append((*t)[:index], (*t)[index+1:]...)
//	}
//
//	func (t *Transactions) Update(id uuid.UUID, description string, accountID uuid.UUID, amount float64, category Category, createdAt time.Time) *Transaction {
//		index := t.indexOf(id)
//		if index == -1 {
//			return nil
//		}
//		transaction := (*t)[index]
//		transaction.Update(description, accountID, amount, category, createdAt)
//
//		return transaction
//	}
func (a *Accounts) Search(search string) []*Account {
	list := make([]*Account, 0)
	for _, account := range *a {
		if strings.Contains(strings.ToLower(account.Name), strings.ToLower(search)) {
			list = append(list, account)
		}
	}
	return list
}

//func (a *Accounts) All() []*Account {
//	list := make([]*Account, len(*a))
//	copy(list, *a)
//	return list
//}

func (a *Accounts) Get(id uuid.UUID) *Account {
	index := a.indexOf(id)
	if index == -1 {
		return nil
	}
	return (*a)[index]
}

//	func (t *Transactions) Reorder(ids []uuid.UUID) []*Transaction {
//		newTransactions := make([]*Transaction, len(ids))
//		for i, id := range ids {
//			newTransactions[i] = (*t)[t.indexOf(id)]
//		}
//		copy(*t, newTransactions)
//		return newTransactions
//	}
//
//	func (t *Transactions) AggregateAmountsByCategory(search string) map[Category]float64 {
//		sums := make(map[Category]float64)
//
//		for _, t := range t.Search(search) {
//			sums[t.Category] += t.Amount
//		}
//		return sums
//	}
//
// //func (t *Transactions) Reorder() {
// //	sort.Slice(*t, func(i, j int) bool {
// //		return (*t)[i].TransactionAt.Before((*t)[j].TransactionAt)
// //	})
// //}
func (a *Accounts) indexOf(id uuid.UUID) int {
	for i, account := range *a {
		if account.ID == id {
			return i
		}
	}
	return -1
}
