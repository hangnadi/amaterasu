package main

import (
	"fmt"
	"github.com/hangnadi/amaterasu/login"
	"github.com/hangnadi/amaterasu/register"
	"net/http"
)

const (
	Address string = ":8080"
)

func main() {
	providers := []string{"email", "google", "facebook"}

	for _, provider := range providers {
		http.HandleFunc(fmt.Sprintf("/auth/%s/login", provider), loginHandler(provider))
		http.HandleFunc(fmt.Sprintf("/auth/%s/callback", provider), callBackHandler(provider))
	}
	http.HandleFunc("/register", registerHandler())
	http.ListenAndServe(Address, nil)
}

func loginHandler(providerName string) http.HandlerFunc {
	return login.ProvideLogin(providerName)
}

func registerHandler() http.HandlerFunc {
	return register.ProvideRegister()
}

func callBackHandler(providerName string) http.HandlerFunc {
	return login.ProvideCallBack(providerName)
}
