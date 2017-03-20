package main

import (
	"log"
	"net/http"
	"time"
)

const (
	Address string = ":8080"
)

func main() {
	router := NewRouter()

	s := &http.Server{
		Addr:           Address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}
	log.Fatal(s.ListenAndServe())
	// providers := []string{"email", "google", "facebook"}

	// for _, provider := range providers {
	// 	http.HandleFunc(fmt.Sprintf("/auth/%s/login", provider), loginHandler(provider))
	// 	http.HandleFunc(fmt.Sprintf("/auth/%s/callback", provider), callBackHandler(provider))
	// }
	// http.HandleFunc("/register", registerHandler())
	// http.ListenAndServe(Address, nil)
}

// func loginHandler(providerName string) http.HandlerFunc {
// 	return login.ProvideLogin(providerName)
// }

// func registerHandler() http.HandlerFunc {
// 	return register.ProvideRegister()
// }

// func callBackHandler(providerName string) http.HandlerFunc {
// 	return login.ProvideCallBack(providerName)
// }
