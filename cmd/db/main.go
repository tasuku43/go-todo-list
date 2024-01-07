package db

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tasuku43/go-todo-list/models"
)

func main() {
	db := models.GormConnect()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Task{})
}
