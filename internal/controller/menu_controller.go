package controller

import (
	"errors"
	"io"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/service"
)

func explainMenu(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		codec.Failure(w, dto.NewApiError("INVALID_IMAGE", "Failed to read image", http.StatusBadRequest))
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		codec.Failure(w, dto.NewApiError("IMAGE_READ_FAILED", "Could not read image data", http.StatusBadRequest))
		return
	}

	format, err := util.DetectImageFormat(imageBytes)
	if err != nil {
		codec.Failure(w, dto.NewApiError("INVALID_IMAGE_FORMAT", "Unsupported or corrupt image", http.StatusBadRequest))
		return
	}
	if format != "jpeg" && format != "jpg" && format != "png" {
		codec.Failure(w, dto.NewApiError("UNSUPPORTED_FORMAT", "Only JPEG, JPG, PNG images are allowed", http.StatusUnsupportedMediaType))
		return
	}

	var fail *fail.Fail
	explanation, err := service.ExplainMenu(imageBytes, format)
	if errors.As(err, &fail) {
		codec.FailureFromFail(w, fail)
		return
	}

	codec.Success(w, http.StatusOK, dto.MenuExplanationResponse{Explanation: explanation})
}
