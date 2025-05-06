package repository

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	dto "github.com/yonsei-autopilot/smart-menu-backend/internal/dto/menu"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

func SearchMenuImage(request *dto.MenuExplanationRequest) (*dto.ImageSearchResult, *fail.Fail) {
	adjustedQuery := url.QueryEscape(request.Name) + "+food+dish+photo+-menu+-logo"

	result, fail := search(adjustedQuery, 3)
	if fail != nil {
		return nil, fail
	}

	return result, nil
}

func search(adjustedQuery string, count int) (*dto.ImageSearchResult, *fail.Fail) {
	searchURL := fmt.Sprintf(
		"https://www.googleapis.com/customsearch/v1?q=%s&cx=%s&key=%s&searchType=image&num=%d",
		adjustedQuery, util.GoogleSearchEngineId, util.GoogleApiKey, count,
	)

	response, err := http.Get(searchURL)
	if err != nil {
		return nil, &fail.GoogleSearchNotWorking
	}
	defer response.Body.Close()

	imageSearchResult, err := codec.DecodeRes[dto.ImageSearchResult](response)
	if err != nil {
		return nil, &fail.InvalidJsonBody
	}

	if failure := imageSearchResult.Validate(); failure != nil {
		return nil, &fail.ResponseValidationFailed
	}

	return imageSearchResult, nil
}
