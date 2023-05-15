package categories

type CreateCategoryRequest struct {
	Name       string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type:text" form:"description"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	
}