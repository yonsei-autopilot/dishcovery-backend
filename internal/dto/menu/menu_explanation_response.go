package dto

type MenuExplanationResponse struct {
	Links []string `json:"links"`
}

func NewMenuExplanationResponse(result *ImageSearchResult) *MenuExplanationResponse {
	links := make([]string, len(result.Items))
	for i, item := range result.Items {
		links[i] = item.Link
	}
	return &MenuExplanationResponse{Links: links}
}
