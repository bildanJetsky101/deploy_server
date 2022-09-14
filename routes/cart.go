package routes

import (
	"server_wb/handlers"
	"server_wb/pkg/middleware"
	"server_wb/pkg/mysql"
	"server_wb/repositories"

	"github.com/gorilla/mux"
)

// Midpss-10

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCarts).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.GetCart)).Methods("GET")
	r.HandleFunc("/user-cart", middleware.Auth(h.GetCarts)).Methods("GET")
	r.HandleFunc("/addCart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
	r.HandleFunc("/clean-cart", middleware.Auth(h.CleaningCart)).Methods("DELETE")
}
