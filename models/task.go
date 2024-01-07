package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name        string    `json:"name"`
	Deadline    time.Time `json:"deadline"`
	Description string    `json:"description"`
}
