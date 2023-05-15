package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"qurban-yuk/dto"
	"qurban-yuk/dto/categories"
	"qurban-yuk/models"
	"qurban-yuk/repositories"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)


type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
}

func HandlerCategory(categoryRepository repositories.CategoryRepository) *handlerCategory {
	return &handlerCategory{categoryRepository}
}

func (h *handlerCategory) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category, err := h.CategoryRepository.GetCategory()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	for i, p := range category {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		category[i].Image = imagePath
	}

	response, _ := json.Marshal(category)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *handlerCategory) GetCategoryID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	category, err := h.CategoryRepository.GetCategoryID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	category.Image = os.Getenv("PATH_FILE") + category.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: category}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataUpload := r.Context().Value("dataFile")
	filename := dataUpload.(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filename, uploader.UploadParams{Folder: "qurban"})

	price, _ := strconv.Atoi(r.FormValue("price"))

	Field := models.Category{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Image:       resp.SecureURL,
	}

	category, err := h.CategoryRepository.CreateCategory(Field)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: category}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataUpload := r.Context().Value("dataFile")
	filename := dataUpload.(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filename, uploader.UploadParams{Folder: "qurban"})


	price, _ := strconv.Atoi(r.FormValue("price"))

	request := categories.CreateCategoryRequest{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Image: resp.SecureURL,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	category := models.Category{}

	category.ID = id

	if request.Name != "" {
		category.Name = request.Name
	}

	if request.Price != 0 {
		category.Price = request.Price
	}

	if filename != "" {
		category.Image = request.Image
	}

	if request.Description != "" {
		category.Description = request.Description
	}

	data, err := h.CategoryRepository.UpdateCategory(category, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	category := models.Category{}

	delCategory, err := h.CategoryRepository.DeleteCategory(category, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: delCategory}
	json.NewEncoder(w).Encode(response)
}
