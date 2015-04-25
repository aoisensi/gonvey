package main

import (
	"net/http"

	"github.com/gorilla/mux"

	ctrl "github.com/aoisensi/gonvey/controller"
)

func Route(serve *http.ServeMux) {
	r := mux.NewRouter()
	r.HandleFunc("/", ctrl.IndexGet).Methods("GET")
	r.HandleFunc("/signup", ctrl.SignupPost).Methods("POST")
	r.HandleFunc("/username_available", ctrl.UsernameAvailableGet).Methods("GET")
	r.HandleFunc("/token", ctrl.TokenPost).Methods("POST")
	p := http.FileServer(http.Dir("./public"))
	serve.Handle("/public/", http.StripPrefix("/public/", p))
	serve.Handle("/", r)
}
