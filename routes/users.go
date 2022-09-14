package routes

import (
	"server_wb/handlers"
	"server_wb/pkg/middleware"
	"server_wb/pkg/mysql"
	"server_wb/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/get-user", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
}
