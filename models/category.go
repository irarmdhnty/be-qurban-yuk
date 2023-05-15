package models

type Category struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
