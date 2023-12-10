package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"spendings/internal/domain"
	"time"
)

func accounts(db *gorm.DB) {
	db.AutoMigrate(&domain.Account{})
	db.Create(domain.NewAccount("Revolut Pro CHF", domain.CHF, 0))
}

func transactions(db *gorm.DB) {
	db.AutoMigrate(&domain.Account{})
	var a1 domain.Account
	db.First(&a1)

	db.AutoMigrate(&domain.Transaction{})
	db.Create(domain.NewTransaction("Car Loan", &a1, 300, domain.LOAN, time.Date(2023, time.November, 20, 0, 0, 0, 0, time.UTC)))
	db.Create(domain.NewTransaction("Electricity", &a1, 29.01, domain.UTILITIES, time.Date(2023, time.November, 16, 0, 0, 0, 0, time.UTC)))
	db.Create(domain.NewTransaction("Badminton", &a1, 13, domain.SPORTS, time.Date(2023, time.November, 10, 0, 0, 0, 0, time.UTC)))
	db.Create(domain.NewTransaction("Apple Cloud", &a1, 0.99, domain.SUBSCRIPTION, time.Date(2023, time.November, 16, 0, 0, 0, 0, time.UTC)))
	db.Create(domain.NewTransaction("Internet", &a1, 28.99, domain.UTILITIES, time.Date(2023, time.November, 16, 0, 0, 0, 0, time.UTC)))
	db.Create(domain.NewTransaction("Migros", &a1, 6.6, domain.GROCERIES, time.Date(2023, time.November, 16, 0, 0, 0, 0, time.UTC)))
	db.Create(domain.NewTransaction("Suan Long", &a1, 10.5, domain.RESTAURANT, time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC)))
	db.Create(domain.NewTransaction("Coop", &a1, 4.7, domain.GROCERIES, time.Date(2023, time.November, 14, 0, 0, 0, 0, time.UTC)))
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	accounts(db)
	transactions(db)
}
