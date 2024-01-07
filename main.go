package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tasuku43/go-todo-list/models"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	dbConn := initDB() // データベース接続の初期化
	defer dbConn.Close()

	r.POST("/tasks", func(c *gin.Context) {
		var task models.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dbConn.Create(&task)
		c.JSON(http.StatusCreated, task)
	})

	r.GET("/tasks", func(c *gin.Context) {
		var tasks []models.Task
		dbConn.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		if err := dbConn.First(&task, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		if err := dbConn.First(&task, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task.UpdatedAt = time.Now()
		dbConn.Save(&task)
		c.JSON(http.StatusOK, task)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		if err := dbConn.First(&task, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		dbConn.Delete(&task)
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})

	r.Run(":8080")
}

func initDB() *gorm.DB {
	connect := models.GormConnect()
	connect.LogMode(true)
	return connect
}
