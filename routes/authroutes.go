package routes

import (
	"qurban-yuk/handlers"
	"qurban-yuk/pkg/midleware"
	"qurban-yuk/pkg/mysql"
	"qurban-yuk/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/regist", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	// for cek auth
	r.HandleFunc("/check-auth", midleware.Auth(h.CheckAuth)).Methods("GET")

}
