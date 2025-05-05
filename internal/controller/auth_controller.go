package controller

import (
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/auth"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/service"
)

func googleLogin(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[dto.GoogleLoginRequest](r)
	if err != nil {
		codec.FailureFromFail(w, &fail.InvalidJsonBody)
		return
	}

	if fail := req.Validate(); fail != nil {
		codec.FailureFromFail(w, fail)
		return
	}

	response, fail := service.GoogleLogin(r.Context(), req.AccessToken)
	if fail != nil {
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

	if fail := req.Validate(); fail != nil {
		codec.FailureFromFail(w, fail)
		return
	}

	response, fail := service.SimpleLogin(r.Context(), req)
	if fail != nil {
		codec.FailureFromFail(w, fail)
		return
	}

	codec.Success(w, http.StatusOK, response)
}

func register(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[dto.RegisterRequest](r)
	if err != nil {
		codec.FailureFromFail(w, &fail.InvalidJsonBody)
		return
	}

	if fail := req.Validate(); fail != nil {
		codec.FailureFromFail(w, fail)
		return
	}

	fail := service.Register(r.Context(), req)
	if fail != nil {
		codec.FailureFromFail(w, fail)
		return
	}

	codec.Success(w, http.StatusOK, nil)
}
