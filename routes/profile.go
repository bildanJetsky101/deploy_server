package routes

import (
	"server_wb/handlers"
	"server_wb/pkg/middleware"
	"server_wb/pkg/mysql"
	"server_wb/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profiles", h.FindProfiles).Methods("GET")
	r.HandleFunc("/get-profile", middleware.Auth(h.GetProfile)).Methods("GET")
	r.HandleFunc("/add-profile", middleware.Auth(middleware.UploadFile(h.CreateProfile))).Methods("POST")
}
