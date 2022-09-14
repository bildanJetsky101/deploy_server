package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	ProductRoutes(r)
	CartRoutes(r)
	AuthRoutes(r)
	TransactionRoutes(r)
	ProfileRoutes(r)
}
