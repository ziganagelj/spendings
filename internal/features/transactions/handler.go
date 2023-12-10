package transactions

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"spendings/internal/domain"
	"spendings/internal/features/accounts"
	"spendings/internal/templates/pages"
	"spendings/internal/templates/partials"
	"spendings/internal/util"
	"strconv"
	"time"
)

type (
	Handler interface {
		//Search : GET /transactions
		Search(w http.ResponseWriter, r *http.Request)
		//Dev : GET /transactions/dev
		Dev(w http.ResponseWriter, r *http.Request)
		// Create : POST /transactions
		Create(w http.ResponseWriter, r *http.Request)
		// Update : PATCH /transactions/{transactionId}
		// Update : POST /transactions/{transactionId}/edit
		Update(w http.ResponseWriter, r *http.Request)
		//// Get : GET /transactions/{transactionId}
		Get(w http.ResponseWriter, r *http.Request)
		// Delete : DELETE /transactions/{transactionId}
		// Delete : POST /transactions/{transactionId}/delete
		Delete(w http.ResponseWriter, r *http.Request)
		// Sort : POST /transactions/sort
		Sort(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		service        Service
		accountService accounts.Service
	}
)

func NewHandler(service Service, accountService accounts.Service) Handler {
	return &handler{
		service:        service,
		accountService: accountService,
	}
}

func Mount(r chi.Router, h Handler) {
	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", h.Search)
		r.Get("/dev", h.Dev)
		r.Post("/", h.Create)
		r.Route("/{transactionId}", func(r chi.Router) {
			r.Patch("/", h.Update)
			r.Post("/edit", h.Update)
			r.Get("/", h.Get)
			r.Delete("/", h.Delete)
			r.Post("/delete", h.Delete)
		})
		r.Post("/sort", h.Sort)
	})
}
func (h handler) Dev(w http.ResponseWriter, r *http.Request) {
	var search = r.URL.Query().Get("search")
	month, _ := strconv.Atoi(r.URL.Query().Get("month"))
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	transactions, err := h.service.Search(r.Context(), search, month, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	aggregate, err := h.service.AggregateAmountsByCategoryAndMonth(r.Context(), search, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accounts, err := h.accountService.Search(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	println(transactions, aggregate, accounts)

	switch util.IsHTMX(r) {
	//case true:
	//	err = partials.RenderTransactions(transactions).Render(r.Context(), w)
	default:
		err = pages.ChartPage(aggregate, "").Render(r.Context(), w)
		//err = partials.AddTransactionForm().Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		err           error
		description   string
		amount        float64
		category      domain.Category
		accountID     uuid.UUID
		account       *domain.Account
		transactionAt time.Time
		transaction   *domain.Transaction
	)

	if err = r.ParseForm(); err != nil {
		util.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Retrieve form values
	description = r.Form.Get("description")
	amountStr := r.Form.Get("amount")
	categoryStr := r.Form.Get("category")
	accountIdStr := r.Form.Get("accountId")
	transactionAtStr := r.Form.Get("transactionAt")

	// Process form values and handle errors
	if amount, err = strconv.ParseFloat(amountStr, 64); err != nil {
		util.RespondWithError(w, "Invalid format for amount", http.StatusBadRequest)
		return
	}
	if category, err = domain.ParseCategory(categoryStr); err != nil {
		util.RespondWithError(w, "Invalid category", http.StatusBadRequest)
		return
	}
	if accountID, err = uuid.Parse(accountIdStr); err != nil {
		util.RespondWithError(w, "Invalid format for accountId", http.StatusBadRequest)
		return
	}
	if account, err = h.accountService.Get(r.Context(), accountID); err != nil {
		util.RespondWithError(w, "Failed to retrieve account", http.StatusBadRequest)
		return
	}
	if transactionAt, err = time.Parse("2006-01-02", transactionAtStr); err != nil {
		util.RespondWithError(w, "Invalid format for transaction date", http.StatusBadRequest)
		return
	}

	// Create transaction and handle error
	if transaction, err = h.service.Add(r.Context(), description, account, amount, category, transactionAt); err != nil {
		util.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch util.IsHTMX(r) {
	case true:
		err = partials.RenderTransaction(transaction).Render(r.Context(), w)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Search(w http.ResponseWriter, r *http.Request) {
	var search = r.URL.Query().Get("search")

	// Retrieve the current year and month
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()

	// Get month and year from the request, use current month/year if not provided
	month, err := strconv.Atoi(r.URL.Query().Get("month"))
	if err != nil {
		month = int(currentMonth)
	}
	year, err := strconv.Atoi(r.URL.Query().Get("year"))
	if err != nil {
		year = currentYear
	}
	transactions, err := h.service.Search(r.Context(), search, month, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accs, err := h.accountService.Search(r.Context(), "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch util.IsHTMX(r) {
	case true:
		err = partials.RenderTransactions(transactions).Render(r.Context(), w)
	default:
		err = pages.TransactionsPage(transactions, accs, domain.GetAllCategories(), search, month, year).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "transactionId")
	var (
		err           error
		transactionID uuid.UUID
		description   string
		amount        float64
		category      domain.Category
		accountID     uuid.UUID
		account       *domain.Account
		transactionAt time.Time
		transaction   *domain.Transaction
	)

	if err = r.ParseForm(); err != nil {
		util.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve form request
	if transactionID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve form values
	description = r.Form.Get("description")
	amountStr := r.Form.Get("amount")
	categoryStr := r.Form.Get("category")
	accountIdStr := r.Form.Get("accountId")
	transactionAtStr := r.Form.Get("transactionAt")

	if amount, err = strconv.ParseFloat(amountStr, 64); err != nil {
		util.RespondWithError(w, "Invalid format for amount", http.StatusBadRequest)
		return
	}
	if category, err = domain.ParseCategory(categoryStr); err != nil {
		util.RespondWithError(w, "Invalid category", http.StatusBadRequest)
		return
	}
	if accountID, err = uuid.Parse(accountIdStr); err != nil {
		util.RespondWithError(w, "Invalid format for accountId", http.StatusBadRequest)
		return
	}
	if account, err = h.accountService.Get(r.Context(), accountID); err != nil {
		util.RespondWithError(w, "Failed to retrieve account", http.StatusBadRequest)
		return
	}
	if transactionAt, err = time.Parse("2006-01-02", transactionAtStr); err != nil {
		util.RespondWithError(w, "Invalid format for transaction date", http.StatusBadRequest)
		return
	}

	if transaction, err = h.service.Update(r.Context(), transactionID, description, account, amount, category, transactionAt); err != nil {
		util.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch util.IsHTMX(r) {
	case true:
		err = partials.RenderTransaction(transaction).Render(r.Context(), w)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "transactionId")
	var transactionID uuid.UUID
	var err error
	if transactionID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Remove(r.Context(), transactionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch util.IsHTMX(r) {
	case true:
		_, err = w.Write([]byte(""))
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "transactionId")
	var transactionID uuid.UUID
	var err error
	if transactionID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction, err := h.service.Get(r.Context(), transactionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch util.IsHTMX(r) {
	case true:
		err = partials.EditTransactionForm(transaction, domain.GetAllCategories()).Render(r.Context(), w)
	default:
		err = pages.TransactionPage(transaction).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Sort(w http.ResponseWriter, r *http.Request) {
	var transactionIDs []uuid.UUID
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, id := range r.Form["id"] {
		var transactionID uuid.UUID
		var err error
		if transactionID, err = uuid.Parse(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		transactionIDs = append(transactionIDs, transactionID)
	}
	if err := h.service.Sort(r.Context(), transactionIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch util.IsHTMX(r) {
	case true:
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
