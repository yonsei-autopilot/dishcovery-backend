package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	GoogleApiKey         string
	GoogleSearchEngineId string
	GeminiApiKey         string
	JwtSecretKey         string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found")
	}

	initializeValues()
}

func initializeValues() {
	key, err := getEnv("GOOGLE_API_KEY")
	if err != nil {
		panic("missing google api key")
	}
	GoogleApiKey = key

	engineId, err := getEnv("GOOGLE_SEARCH_ENGINE_ID")
	if err != nil {
		panic("missing google search engine id")
	}
	GoogleSearchEngineId = engineId

	apiKey, err := getEnv("GEMINI_API_KEY")
	if err != nil {
		panic("missing Gemini API key")
	}
	GeminiApiKey = apiKey

	secretKey, err := getEnv("JWT_SECRET_KEY")
	if err != nil {
		panic("missing jwt secret key")
	}
	JwtSecretKey = secretKey
}

func getEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variable '%s' not found", key)
	}
	return value, nil
}
