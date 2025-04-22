package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/service"
)

type responseBody struct {
	Explanation string `json:"explanation,omitempty"`
	Error       string `json:"error,omitempty"`
}

func uploadMenuImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	file, _, err := r.FormFile("image")
	if err != nil {
		writeJsonError(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		writeJsonError(w, "Could not read image data", http.StatusInternalServerError)
		return
	}

	response, err := service.SendToGemini(imageBytes)
	if err != nil {
		writeJsonError(w, "Failed to process image with Gemini", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(responseBody{
		Explanation: response,
	})
}

func writeJsonError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responseBody{
		Error: message,
	})
}
