package controller

import (
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	contextHelper "github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/context_helper"
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
