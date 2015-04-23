package controller

import (
	"net/http"

	"github.com/aoisensi/gonvey/models"
	"github.com/k0kubun/pp"
)

type SignUpResult struct {
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	u := models.User{Username: username, Password: password}
	pp.Println(u)
	if !u.Check() {
		w.Write(MakeErrorResult(ErrorBadUsername))
		return
	}
	ok := DB.NewRecord(u)
	if !ok {
		w.Write(MakeErrorResult(ErrorExistUsername))
		return
	}
	pp.Println(u)
	DB.Create(&u)
	w.Write(MakeJsonResult(ok))
}
