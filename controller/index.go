package controller

import (
	"io"
	"net/http"
	"os"
)

func IndexGet(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("view/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
