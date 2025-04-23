package controller

import (
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/encoder"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
)

func checkHealth(w http.ResponseWriter, r *http.Request) {
	encoder.Success(w, http.StatusOK, dto.HealthResponse{Description: "Server is healthy"})
}
