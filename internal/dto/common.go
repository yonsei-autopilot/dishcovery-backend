package dto

import "github.com/yonsei-autopilot/smart-menu-backend/internal/fail"

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func NewApiError(code string, message string, status int) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

func (e *ApiError) Error() string {
	return e.Message
}

func From(fail fail.Fail) *ApiError {
	return &ApiError{
		Code:    fail.Code,
		Message: fail.Message,
		Status:  fail.Status,
	}
}

type ApiResponse struct {
	IsSuccess bool        `json:"isSuccess"`
	Data      interface{} `json:"data,omitempty"`
	Error     *ApiError   `json:"error,omitempty"`
}
