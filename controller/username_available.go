package controller

import (
	"net/http"

	"github.com/aoisensi/gonvey/models"
)

type UsernameAvailableState struct {
	Creatable bool   `json:"creatable"`
	Username  string `json:"username"`
	Reason    string `json:"reason"`
}

func UsernameAvailableGet(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	state := UsernameAvailableState{Username: username}
	parse := models.CheckUsername(username)
	if !parse {
		state.Reason = "Invalid character or Invalid chara count"
		w.Write(MakeJsonResult(state))
		return
	}
	found := !DB.First(&models.User{}, "username = ?", username).RecordNotFound()
	if found {
		state.Reason = "This username is already used"
		w.Write(MakeJsonResult(state))
		return
	}
	state.Creatable = true
	w.Write(MakeJsonResult(state))
	return
}
