package controller

import (
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	contextHelper "github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/context_helper"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/user"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/service"
)

func getDislikeFoods(w http.ResponseWriter, r *http.Request) {
	id, fail := contextHelper.GetUserId(r.Context())
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	response, fail := service.GetDislikeFoods(r.Context(), id)
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	codec.Success(w, http.StatusOK, response)
}

func updateDislikeFoods(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[dto.UpdateDislikeFoodsResponse](r)
	if err != nil {
		codec.Failure(w, &fail.InvalidJsonBody)
		return
	}

	id, fail := contextHelper.GetUserId(r.Context())
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	fail = service.UpdateDislikeFoods(r.Context(), id, req)
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	codec.Success(w, http.StatusOK, nil)
}
