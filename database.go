package main

import (
	"log"

	"github.com/jinzhu/gorm"

	ctrl "github.com/aoisensi/gonvey/controller"
	"github.com/aoisensi/gonvey/models"
	_ "github.com/mattn/go-sqlite3"
)

var DB gorm.DB

func InitDB() {
	db, err := gorm.Open("sqlite3", "gonvey.db")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
	ctrl.DB = DB
}
