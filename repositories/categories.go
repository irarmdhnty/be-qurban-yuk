package repositories

import (
	"gorm.io/gorm"
	"qurban-yuk/models"
)

type CategoryRepository interface{
	GetCategory() ([]models.Category, error)
	GetCategoryID(ID int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category, ID int) (models.Category, error)
	DeleteCategory(category models.Category, ID int) (models.Category, error)
} 

func RepositoryCategories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCategory() ([]models.Category, error) {
	var category []models.Category
	err := r.db.Find(&category).Error
	return category, err
}

func (r *repository) GetCategoryID(ID int) (models.Category, error){
	var categoryId models.Category

	err := r.db.First(&categoryId, ID).Error
	return categoryId, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error){
	err := r.db.Create(&category).Error
	return category, err
}

func (r *repository) UpdateCategory(category models.Category, ID int) (models.Category, error){
	err := r.db.Model(&category).Where("id =?", ID).Updates(&category).Error
	return category, err
}

func (r * repository) DeleteCategory(category models.Category, ID int) (models.Category, error){
	err := r.db.Delete(&category, ID).Error
	return category, err
}