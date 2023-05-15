package databases

import (
	"fmt"
	"qurban-yuk/models"
	"qurban-yuk/pkg/mysql"
)

func RunMigrate() {
	if err := mysql.DB.AutoMigrate(&models.User{}, &models.Category{}); err != nil {
		fmt.Println("err")
		panic("Migrate failed")
	}
	fmt.Println("Migrate Successful")
}
