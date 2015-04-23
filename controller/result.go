package controller

import "encoding/json"

var (
	ErrorSuccess        = Error{Code: 0x0000, Message: ""}
	ErrorInternalServer = Error{Code: 0x0001, Message: "Internal server error."}
	ErrorExistUsername  = Error{Code: 0x0101, Message: "The username is already used."}
	ErrorWeekPassword   = Error{Code: 0x0102, Message: "This password is so much week."}
	ErrorBadUsername    = Error{Code: 0x0103, Message: "This username can not be used."}
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Result struct {
	IsError bool        `json:"is_error"`
	Error   Error       `json:"error"`
	State   interface{} `json:"state"`
}

func MakeErrorResult(e Error) []byte {
	result := Result{IsError: true, Error: e, State: nil}
	body, _ := json.Marshal(&result)
	return body
}

func MakeJsonResult(state interface{}) []byte {
	result := Result{IsError: false, Error: ErrorSuccess, State: state}
	body, _ := json.Marshal(&result)
	return body
}
