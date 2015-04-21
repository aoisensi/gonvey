package main

import (
	"net/http"

	"github.com/aoisensi/gonvey/models"
)

func SignupPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	u := models.User{Name: username, Password: password}
	DB.NewRecord(u)
	DB.Create(&u)
}
