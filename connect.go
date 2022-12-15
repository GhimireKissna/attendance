package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() (*gorm.DB, error) {
	dsn := "root:Krishna@7@tcp(127.0.0.1:3306)/attendance?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db, err
}
