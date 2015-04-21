package main

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/aoisensi/gonvey/models"
)

var DB gorm.DB

func InitDB() {
	db, err := gorm.Open("sqlite3", "gonvey.go")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
