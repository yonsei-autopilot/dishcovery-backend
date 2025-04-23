package dto

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

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ApiError   `json:"error,omitempty"`
}
