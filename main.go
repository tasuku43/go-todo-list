package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tasuku43/go-todo-list/models"
	"github.com/tasuku43/go-todo-list/pkg/presentation/rest/handlers"
)

func main() {
	r := gin.Default()
	dbConn := initDB() // データベース接続の初期化
	defer dbConn.Close()
	taskHandler := handlers.NewTaskHandler(dbConn)

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	r.Run(":8080")
}

func initDB() *gorm.DB {
	connect := models.GormConnect()
	connect.LogMode(true)
	return connect
}
