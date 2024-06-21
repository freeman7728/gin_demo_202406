package dto

const (
	SUCCESS = 0
	ERROR   = 500
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
