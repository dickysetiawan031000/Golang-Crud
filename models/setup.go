package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_restapi_gin"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	DB = db
}
