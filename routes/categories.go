package routes

import (
	"qurban-yuk/handlers"
	"qurban-yuk/pkg/midleware"
	"qurban-yuk/pkg/mysql"
	"qurban-yuk/repositories"

	"github.com/gorilla/mux"
)

func CategoriesRoute(r *mux.Router) {
	categoryRepo := repositories.RepositoryCategories(mysql.DB)
	h := handlers.HandlerCategory(categoryRepo)

	r.HandleFunc("/categories", h.GetCategory).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategoryID).Methods("GET")
	r.HandleFunc("/category/create", midleware.Auth(midleware.UploadFile(h.CreateCategory))).Methods("POST")
	r.HandleFunc("/category/update/{id}", midleware.Auth(midleware.UploadFile(h.UpdateCategory))).Methods("PATCH")
	r.HandleFunc("/category/delete/{id}", h.DeleteCategory).Methods("DELETE")

}
