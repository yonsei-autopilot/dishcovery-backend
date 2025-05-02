package repository

import (
	"fmt"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/dto"
)

func FetchUserInfo(accessToken string) (*dto.UserInfoResponse, error) {
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("userinfo API error: %s", resp.Status)
	}

	userInfo, err := codec.DecodeRes[dto.UserInfoResponse](resp)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
