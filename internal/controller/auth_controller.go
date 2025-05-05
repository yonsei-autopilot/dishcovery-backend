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
	req, err := codec.DecodeReq[dto.GoogleLoginRequest](r)
	if err != nil {
		codec.FailureFromFail(w, &fail.InvalidJsonBody)
		return
	}

	if req.AccessToken == "" {
		codec.FailureFromFail(w, &fail.RequestValidationFailed)
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

func simpleLogin(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[dto.SimpleLoginRequest](r)
	if err != nil {
		codec.FailureFromFail(w, &fail.InvalidJsonBody)
		return
	}

	if req.LoginId == "" || req.Password == "" {
		codec.FailureFromFail(w, &fail.RequestValidationFailed)
		return
	}

	response, err := service.SimpleLogin(r.Context(), req)
	var fail *fail.Fail
	if errors.As(err, &fail) {
		codec.FailureFromFail(w, fail)
		return
	}

	codec.Success(w, http.StatusOK, response)
}
