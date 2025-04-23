package controller

import (
	"io"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/service"
)

func explainMenu(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		util.Error(w, dto.NewApiError("INVALID_IMAGE", "Failed to read image", http.StatusBadRequest))
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		util.Error(w, dto.NewApiError("IMAGE_READ_FAILED", "Could not read image data", http.StatusBadRequest))
		return
	}

	format, err := util.DetectImageFormat(file)
	if err != nil {
		util.Error(w, dto.NewApiError("INVALID_IMAGE_FORMAT", "Unsupported or corrupt image", http.StatusBadRequest))
		return
	}
	if format != "jpeg" && format != "jpg" && format != "png" {
		util.Error(w, dto.NewApiError("UNSUPPORTED_FORMAT", "Only JPEG, JPG, PNG images are allowed", http.StatusUnsupportedMediaType))
		return
	}

	explanation, err := service.ExplainMenu(imageBytes, format)
	if err != nil {
		util.Error(w, dto.NewApiError("GEMINI_PROCESS_FAILED", err.Error(), http.StatusOK))
		return
	}

	util.JSON(w, http.StatusOK, dto.MenuExplanationResponse{Explanation: explanation})
}
