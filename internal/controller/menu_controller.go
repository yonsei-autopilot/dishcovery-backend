package controller

import (
	"io"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	contextHelper "github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/context_helper"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu/google_tts"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/service"
)

func translateMenu(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		codec.Failure(w, &fail.InvalidImage)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		codec.Failure(w, &fail.ImageReadFailed)
		return
	}

	format, err := util.DetectImageFormat(imageBytes)
	if err != nil {
		codec.Failure(w, &fail.InvalidImageFormat)
		return
	}
	if format != "jpeg" && format != "jpg" && format != "png" {
		codec.Failure(w, &fail.UnsupportedImageFormat)
		return
	}

	id, fail := contextHelper.GetUserId(r.Context())
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	menuTranslation, fail := service.TranslateMenu(r.Context(), id, imageBytes, format)
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	codec.Success(w, menuTranslation)
}

func explainMenu(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[dto.MenuExplanationRequest](r)
	if err != nil {
		codec.Failure(w, &fail.InvalidJsonBody)
		return
	}

	if fail := req.Validate(); fail != nil {
		codec.Failure(w, fail)
		return
	}

	id, fail := contextHelper.GetUserId(r.Context())
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	menuExplanation, fail := service.ExplainMenu(r.Context(), id, req)
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	codec.Success(w, menuExplanation)
}

func orderMenu(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[dto.MenuOrderRequest](r)
	if err != nil {
		codec.Failure(w, &fail.InvalidJsonBody)
		return
	}

	if fail := req.Validate(); fail != nil {
		codec.Failure(w, fail)
		return
	}

	id, fail := contextHelper.GetUserId(r.Context())
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	menuOrder, fail := service.OrderMenu(r.Context(), id, req)
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	codec.Success(w, menuOrder)
}

func getForeignLanguageOfMenu(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[google_tts.LanguageCodeForGoogleTtsRequest](r)
	if err != nil {
		codec.Failure(w, &fail.InvalidJsonBody)
		return
	}

	if fail := req.Validate(); fail != nil {
		codec.Failure(w, fail)
		return
	}

	languageCode, fail := service.GetForeignLanguageOfMenu(r.Context(), req)
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	codec.Success(w, languageCode)
}

func getMenuOrderSpeech(w http.ResponseWriter, r *http.Request) {
	req, err := codec.DecodeReq[google_tts.MenuOrderSpeechRequest](r)
	if err != nil {
		codec.Failure(w, &fail.InvalidJsonBody)
		return
	}

	if fail := req.Validate(); fail != nil {
		codec.Failure(w, fail)
		return
	}

	languageCode, fail := service.GetMenuOrderSpeech(r.Context(), req)
	if fail != nil {
		codec.Failure(w, fail)
		return
	}

	codec.Success(w, languageCode)
}
