package model

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
