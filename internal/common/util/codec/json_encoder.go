package codec

import (
	"encoding/json"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
)

func Success(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(dto.ApiResponse{
		Success: true,
		Data:    data,
	})
}

func Failure(w http.ResponseWriter, apiErr *dto.ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.Status)

	json.NewEncoder(w).Encode(dto.ApiResponse{
		Success: false,
		Error:   apiErr,
	})
}
