package service

import (
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/repository"
)

func SearchImage(request *dto.MenuExplanationRequest) (*dto.MenuExplanationResponse, *fail.Fail) {
	response, fail := repository.SearchMenuImage(request)
	if fail != nil {
		return nil, fail
	}

	return response, nil
}
