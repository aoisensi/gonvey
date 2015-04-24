package controller

import (
	"net/http"

	"github.com/aoisensi/gonvey/models"
)

type SignUpResult struct {
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	u := models.User{Username: username, Password: password}
	if !u.Check() {
		w.Write(MakeErrorResult(ErrorBadUsername))
		return
	}
	if !DB.First(&models.User{}, "username = ?", username).RecordNotFound() {
		w.Write(MakeErrorResult(ErrorExistUsername))
		return
	}
	DB.NewRecord(u)
	DB.Create(&u)
	w.Write(MakeJsonResult(true))
}
