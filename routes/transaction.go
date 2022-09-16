package routes

import (
	"server_wb/handlers"
	"server_wb/pkg/middleware"
	"server_wb/pkg/mysql"
	"server_wb/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransaction)).Methods("GET")
	r.HandleFunc("/add-transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	// r.HandleFunc("/product/{id}", h.UpdateProduct).Methods("PATCH")
	// r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
}
