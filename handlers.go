package main

import (
	"encoding/json"
	"fmt"
	"github.com/hangnadi/amaterasu/register"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json; charset=UTF-8"
)

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func LoginEmail(rw http.ResponseWriter, req *http.Request) {

}

func LoginFacebook(rw http.ResponseWriter, req *http.Request) {

}

func CallbackFacebook(rw http.ResponseWriter, req *http.Request) {

}

func LoginGoogle(rw http.ResponseWriter, req *http.Request) {

}

func CallbackGoogle(rw http.ResponseWriter, req *http.Request) {

}

func RegisterEmail(rw http.ResponseWriter, req *http.Request) {

	var responseJson ResponseJSON

	data, err := register.HandleRegister(rw, req)

	jsonData, err1 := json.Marshal(data)

	rw.Header().Set(ContentType, ApplicationJson)
	rw.WriteHeader(http.StatusOK)

	if err1 != nil {
		responseJson = ResponseJSON{
			Status:       "OK",
			MessageError: "internal server error",
		}
	} else {
		if err == nil {
			responseJson = ResponseJSON{
				Status: "OK",
				Data:   string(jsonData),
			}
		} else {
			responseJson = ResponseJSON{
				Status:       "OK",
				MessageError: err.Error(),
			}
		}
	}

	if err := json.NewEncoder(rw).Encode(responseJson); err != nil {
		panic(err)
	}
}
