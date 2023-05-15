package routes

import "github.com/gorilla/mux"

func RoutesInit(r *mux.Router) {
	CategoriesRoute(r)
	AuthRoutes(r)
	UserRoute(r)
}