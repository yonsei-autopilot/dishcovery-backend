package dto

type GetDislikeFoodsResponse struct {
	DislikeFoods []string `json:"dislikeFoods"`
}

func FromDislikeFoods(dislikeFoods []string) *GetDislikeFoodsResponse {
	return &GetDislikeFoodsResponse{DislikeFoods: dislikeFoods}
}
