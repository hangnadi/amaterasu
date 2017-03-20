package main

type ResponseJSON struct {
	Status       string `json:"status"`
	MessageError string `json:"message_error, omitempty"`
	Data         string `json:"data, omitempty"`
}
