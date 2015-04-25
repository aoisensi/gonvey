package controller

import (
	"log"
	"net/http"

	"github.com/RangelReale/osin"
)

func TokenPost(w http.ResponseWriter, r *http.Request) {
	resp := OAuth.NewResponse()
	defer resp.Close()

	if ar := OAuth.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		OAuth.FinishAccessRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		log.Println("Error: ", resp.InternalError)
	}
	osin.OutputJSON(resp, w, r)
}
