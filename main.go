package main

import (
	"fmt"
	"net/http"
	"os"
	"server_wb/database"
	"server_wb/pkg/mysql"
	"server_wb/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	//var port = os.Getenv("PORT")

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	fmt.Println(os.Getenv("DB_PORT"))

	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	// var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	// var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:" + os.Getenv("PORT"))
	// Embed the setup allowed in 2 parameter on this below code ...
	//http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(r))

}
