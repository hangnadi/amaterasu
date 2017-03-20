package register

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidUsername  = errors.New("invalid username")
	ErrInvalidFullname  = errors.New("invalid fullname")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrRequiredPassword = errors.New("password required")
	ErrInvalidPassword  = errors.New("password must between 8 and 16 character")
	ErrInternalServer   = errors.New("internal server error")
)

func HandleRegister(rw http.ResponseWriter, req *http.Request) (ResponseModel, error) {

	responseModel := ResponseModel{
		success: 0,
	}

	requestModel := RequestModel{
		username: req.PostFormValue("username"),
		fullname: req.PostFormValue("fullname"),
		email:    req.PostFormValue("email"),
		password: req.PostFormValue("password"),
	}

	if requestModel.username == "" {
		return responseModel, ErrInvalidUsername
	}

	if requestModel.fullname == "" {
		return responseModel, ErrInvalidFullname
	}

	if requestModel.email == "" {
		return responseModel, ErrInvalidEmail
	}

	if requestModel.password == "" {
		return responseModel, ErrRequiredPassword
	}

	if len(requestModel.password) < 8 || len(requestModel.password) > 16 {
		return responseModel, ErrInvalidPassword
	}

	status := NewAccounts(&requestModel)
	responseModel.success = status

	if status != 1 {
		return responseModel, ErrInternalServer
	} else {
		return responseModel, nil
	}

}
