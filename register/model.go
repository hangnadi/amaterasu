package register

type RequestModel struct {
	username     string
	fullname     string
	email        string
	password     string
	passwordHash string
}

type ResponseModel struct {
	success int `json:"is_success"`
}
