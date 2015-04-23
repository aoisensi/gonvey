package main

import (
	"flag"
	"log"
	"net/http"
)

var FlagConfig string

func init() {
	flag.StringVar(&FlagConfig, "config", "gonvey.toml", "Config file path")
	InitDB()
}

func main() {
	flag.Parse()
	if err := LoadConfig(); err != nil {
		log.Fatal(err)
	}
	serve := http.DefaultServeMux
	Route(serve)
	log.Println("Start http server.")
	http.ListenAndServe(":8000", serve)
}
