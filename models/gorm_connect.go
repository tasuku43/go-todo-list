package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_todo_list"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	db.LogMode(true)
	return db
}
