package main

import (
	"fmt"
	"net/http"
	"os"
	"qurban-yuk/databases"
	"qurban-yuk/pkg/mysql"
	"qurban-yuk/routes"
	"github.com/joho/godotenv"
	"log"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	databases.RunMigrate()
	r := mux.NewRouter()

	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	// uploads path prefix
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	//env
	if err := godotenv.Load(); err != nil {
		log.Println(" No ENV file found")
	}

	var allowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})
	var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")

	fmt.Println("SERVER Running on Port 8000")
	http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(r))
}