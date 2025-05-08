package dto

type GetLanguageResponse struct {
	Language string `json:"language"`
}

func FromLanguage(language string) *GetLanguageResponse {
	return &GetLanguageResponse{Language: language}
}
