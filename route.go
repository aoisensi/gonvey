package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route(serve *http.ServeMux) {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexGet).Methods("GET")
	r.HandleFunc("/signup", SignupPost).Methods("POST")
	p := http.FileServer(http.Dir("./public"))
	serve.Handle("/public/", http.StripPrefix("/public/", p))
	serve.Handle("/", r)
}
