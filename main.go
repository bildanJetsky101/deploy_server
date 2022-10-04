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
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	fmt.Println(os.Getenv("DB_PORT"))

	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	// var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	var port = os.Getenv("PORT")

	fmt.Println("server running localhost:" + port)

	http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), AllowedMethods, AllowedOrigins)(r))

}
