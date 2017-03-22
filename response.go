package main

type ResponseJSON struct {
	Status       string      `json:"status"`
	MessageError string      `json:"message_error, omitempty"`
	Data         interface{} `json:"data, omitempty"`
}
