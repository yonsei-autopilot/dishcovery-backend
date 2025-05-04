package codec

import (
	"encoding/json"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

func Success(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(dto.ApiResponse{
		IsSuccess: true,
		Data:      data,
	})
}

func Failure(w http.ResponseWriter, apiErr *dto.ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.Status)

	json.NewEncoder(w).Encode(dto.ApiResponse{
		IsSuccess: false,
		Error:     apiErr,
	})
}

func FailureFromFail(w http.ResponseWriter, fail *fail.Fail) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(fail.Status)

	apiErr := dto.From(*fail)

	json.NewEncoder(w).Encode(dto.ApiResponse{
		IsSuccess: false,
		Error:     apiErr,
	})
}
