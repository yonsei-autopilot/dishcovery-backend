package controller

import (
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
)

func checkHealth(w http.ResponseWriter, r *http.Request) {
	codec.Success(w, dto.HealthResponse{Description: "Server is healthy"})
}
