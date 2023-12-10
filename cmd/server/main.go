package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"spendings/internal/assets"
	"spendings/internal/domain"
	"spendings/internal/features/accounts"
	"spendings/internal/features/transactions"
	"time"
)

func main() {
	var port = ":3000"

	flag.StringVar(&port, "port", port, "port to listen on")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Fetch all transactions from the database
	var as []domain.Account
	if err := db.Find(&as).Error; err != nil {
		log.Fatal("Failed to fetch transactions:", err)
	}

	var ts []domain.Transaction
	if err := db.Preload("Account").Find(&ts).Error; err != nil {
		log.Fatal("Failed to fetch transactions:", err)
	}

	t := domain.NewTransactions(ts)
	a := domain.NewAccounts(as)

	router := chi.NewRouter()
	transactions.Mount(router, transactions.NewHandler(transactions.NewService(db, t), accounts.NewService(db, a)))
	assets.Mount(router)

	server := &http.Server{
		Addr:    "localhost" + port,
		Handler: http.TimeoutHandler(router, 30*time.Second, "request timed out"),
	}

	fmt.Printf("Listening on http://localhost%s/transactions/dev\n", port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
