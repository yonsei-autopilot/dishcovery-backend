package controller

import (
	"net/http"
)

func RegisterRoutes(handler *http.ServeMux) {
	handler.HandleFunc("GET /health", checkHealth)
	handler.HandleFunc("POST /menus/explanation", explainMenu)
	handler.HandleFunc("GET /menus/explanation/test-page", renderMenuExplanationPage)
}
