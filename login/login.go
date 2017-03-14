package login

import (
	"net/http"
)

func main() {

}

func ProvideLogin(providerName string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}

func ProvideCallBack(providerName string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}
