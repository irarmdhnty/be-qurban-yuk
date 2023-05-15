package repositories

import (
	"qurban-yuk/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() ([]models.User, error)
	GetUserID(ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *repository) GetUserID(ID int) (models.User, error) {
	var UserId models.User

	err := r.db.First(&UserId, ID).Error
	return UserId, err
}
