package models

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"is_success"`
}
