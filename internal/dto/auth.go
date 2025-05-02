package dto

type LoginRequest struct {
	AccessToken string `json:"accessToken"`
}

// TODO - 토큰 반환하게 수정해야 함
type LoginResponse struct {
	Name  string `json:"accessToken"`
	Email string `json:"email"`
}
