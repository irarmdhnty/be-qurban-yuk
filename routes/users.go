package routes

import (
	"qurban-yuk/handlers"
	"qurban-yuk/pkg/mysql"
	"qurban-yuk/repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetuserId).Methods("GET")
}
