package controller

import (
	"errors"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/service"
)

func googleLogin(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[dto.LoginRequest](r)
	if err != nil {
		codec.Failure(w, dto.NewApiError("Invalid JSON body", err.Error(), http.StatusBadRequest))
		return
	}

	if req.AccessToken == "" {
		codec.Failure(w, dto.NewApiError("Missing accessToken", "Missing accessToken", http.StatusBadRequest))
		return
	}

	response, err := service.GoogleLogin(r.Context(), req.AccessToken)
	var fail *fail.Fail
	if errors.As(err, &fail) {
		codec.FailureFromFail(w, fail)
		return
	}

	codec.Success(w, http.StatusOK, response)
}
