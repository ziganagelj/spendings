package accounts

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"spendings/internal/util"
)

type (
	Handler interface {
		//Search : GET /accounts
		Search(w http.ResponseWriter, r *http.Request)
		//Dev : GET /transactions/dev
		//Dev(w http.ResponseWriter, r *http.Request)
		//// Create : POST /transactions
		//Create(w http.ResponseWriter, r *http.Request)
		//// Update : PATCH /transactions/{transactionId}
		//// Update : POST /transactions/{transactionId}/edit
		//Update(w http.ResponseWriter, r *http.Request)
		////// Get : GET /transactions/{transactionId}
		//Get(w http.ResponseWriter, r *http.Request)
		//// Delete : DELETE /transactions/{transactionId}
		//// Delete : POST /transactions/{transactionId}/delete
		//Delete(w http.ResponseWriter, r *http.Request)
		//// Sort : POST /transactions/sort
		//Sort(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		service Service
	}
)

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func Mount(r chi.Router, h Handler) {
	r.Route("/accounts", func(r chi.Router) {
		r.Get("/", h.Search)
	})
}

//func (h handler) Dev(w http.ResponseWriter, r *http.Request) {
//	var search = r.URL.Query().Get("search")
//	transactions, err := h.service.Search(r.Context(), search)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	aggregate, err := h.service.AggregateAmountsByCategory(r.Context(), search)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	println(transactions, aggregate)
//
//	switch util.IsHTMX(r) {
//	//case true:
//	//	err = partials.RenderTransactions(transactions).Render(r.Context(), w)
//	default:
//		err = pages.Home(transactions, search, aggregate).Render(r.Context(), w)
//		//err = partials.AddTransactionForm().Render(r.Context(), w)
//	}
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

//func (h handler) Create(w http.ResponseWriter, r *http.Request) {
//	if err := r.ParseForm(); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	var description = r.Form.Get("description")
//	//amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
//	//if err != nil {
//	//	http.Error(w, "Invalid format for amount", http.StatusBadRequest)
//	//	return
//	//}
//	//createdAt, err := time.Parse("2006-01-02", r.Form.Get("createdAt"))
//	//if err != nil {
//	//	http.Error(w, "Invalid format for createdAt", http.StatusBadRequest)
//	//	return
//	//}
//	a1 := domain.NewAccount("Revolut Pro CHF", domain.CHF, 0)
//	transaction, err := h.service.Add(r.Context(), description, a1.ID, 10, domain.GROCERIES, time.Now())
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	switch util.IsHTMX(r) {
//	case true:
//		err = partials.RenderTransaction(transaction).Render(r.Context(), w)
//	default:
//		http.Redirect(w, r, "/", http.StatusFound)
//	}
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

func (h handler) Search(w http.ResponseWriter, r *http.Request) {
	var search = r.URL.Query().Get("search")
	transactions, err := h.service.Search(r.Context(), search)
	println(transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch util.IsHTMX(r) {
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//err = pages.Home(transactions, search, aggregations).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//func (h handler) Update(w http.ResponseWriter, r *http.Request) {
//	var id = chi.URLParam(r, "transactionId")
//	var transactionID uuid.UUID
//	var accountID uuid.UUID
//	var err error
//	if transactionID, err = uuid.Parse(id); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	if err := r.ParseForm(); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	var accountId = r.Form.Get("accountId")
//	if accountID, err = uuid.Parse(accountId); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	var description = r.Form.Get("description")
//	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
//	if err != nil {
//		// Handle error
//		http.Error(w, "Invalid format for amount", http.StatusBadRequest)
//		return
//	}
//	var category = domain.Category(r.Form.Get("category"))
//	//TODO
//	//createdAt, err := time.Parse("2006-01-02", r.Form.Get("createdAt"))
//	//if err != nil {
//	//	// Handle error
//	//	http.Error(w, "Invalid format for createdAt", http.StatusBadRequest)
//	//	return
//	//}
//
//	transaction, err := h.service.Update(r.Context(), transactionID, description, accountID, amount, category, time.Now())
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	switch util.IsHTMX(r) {
//	case true:
//		err = partials.RenderTransaction(transaction).Render(r.Context(), w)
//	default:
//		http.Redirect(w, r, "/", http.StatusFound)
//	}
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
//	var id = chi.URLParam(r, "transactionId")
//	var transactionID uuid.UUID
//	var err error
//	if transactionID, err = uuid.Parse(id); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if err := h.service.Remove(r.Context(), transactionID); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	switch util.IsHTMX(r) {
//	case true:
//		_, err = w.Write([]byte(""))
//	default:
//		http.Redirect(w, r, "/", http.StatusFound)
//	}
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (h handler) Get(w http.ResponseWriter, r *http.Request) {
//	var id = chi.URLParam(r, "transactionId")
//	var transactionID uuid.UUID
//	var err error
//	if transactionID, err = uuid.Parse(id); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	transaction, err := h.service.Get(r.Context(), transactionID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	switch util.IsHTMX(r) {
//	case true:
//		err = partials.EditTransactionForm(transaction).Render(r.Context(), w)
//	default:
//		err = pages.TransactionPage(transaction).Render(r.Context(), w)
//	}
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (h handler) Sort(w http.ResponseWriter, r *http.Request) {
//	var transactionIDs []uuid.UUID
//	if err := r.ParseForm(); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	for _, id := range r.Form["id"] {
//		var transactionID uuid.UUID
//		var err error
//		if transactionID, err = uuid.Parse(id); err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		transactionIDs = append(transactionIDs, transactionID)
//	}
//	if err := h.service.Sort(r.Context(), transactionIDs); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	switch util.IsHTMX(r) {
//	case true:
//		w.WriteHeader(http.StatusNoContent)
//	default:
//		http.Redirect(w, r, "/", http.StatusFound)
//	}
//}
