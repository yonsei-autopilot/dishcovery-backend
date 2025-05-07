package gemini

import (
	"context"
	"errors"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
	"google.golang.org/api/option"
)

var (
	once   sync.Once
	client *genai.Client
)

type GeminiRequest struct {
	ctx                context.Context
	modelName          string
	prompt             string
	imageBytes         []byte
	imageFormat        string
	isStructuredOutput bool
	responseType       string
	temperature        float32
	topP               float32
	topK               int32
	structObj          any
}

// initialize sets up the Gemini client.
func initialize() {
	ctx := context.Background()

	newClient, err := genai.NewClient(ctx, option.WithAPIKey(util.GeminiApiKey))
	if err != nil {
		panic("failed to create Gemini client")
	}

	client = newClient
}

// InitializeGeminiClient initializes the Gemini client only once using sync.Once.
func InitializeGeminiClient() {
	once.Do(func() {
		initialize()
	})
}

// GeminiRequestBuilder creates GeminiRequest
func GeminiRequestBuilder() *GeminiRequest {
	return &GeminiRequest{
		temperature: 1,
		topP:        0.6,
		topK:        64,
	}
}

func (req *GeminiRequest) WithContext(ctx context.Context) *GeminiRequest {
	req.ctx = ctx
	return req
}

func (req *GeminiRequest) WithModel(modelName string) *GeminiRequest {
	req.modelName = modelName
	return req
}

func (req *GeminiRequest) WithPrompt(prompt string) *GeminiRequest {
	req.prompt = prompt
	return req
}

func (req *GeminiRequest) WithImage(imageBytes []byte, format string) *GeminiRequest {
	req.imageBytes = imageBytes
	req.imageFormat = format
	return req
}

func (req *GeminiRequest) WithTemperature(temp float32) *GeminiRequest {
	req.temperature = temp
	return req
}

func (req *GeminiRequest) WithTopP(topp float32) *GeminiRequest {
	req.topP = topp
	return req
}

func (req *GeminiRequest) WithTopK(topk int32) *GeminiRequest {
	req.topK = topk
	return req
}

func (req *GeminiRequest) ExpectStructuredOutput(obj any) *GeminiRequest {
	req.isStructuredOutput = true
	req.structObj = obj
	return req
}

func (req *GeminiRequest) ExpectResponseType(rtype string) *GeminiRequest {
	req.responseType = rtype
	return req
}

func (req *GeminiRequest) Generate() (string, error) {
	// 클라이언트 초기화 안 되었을 경우 에러
	if client == nil {
		return "", errors.New("gemini client is not initialized")
	}

	// 모델 지정 안 되었을 경우 에러
	if req.modelName == "" {
		return "", errors.New("gemini model name is required")
	}
	model := client.GenerativeModel(req.modelName)

	// Structured Output 모드일 경우 ResponseMIMEType과 Schema 설정
	if req.isStructuredOutput {
		model.ResponseMIMEType = "application/json"
		schema := StructToSchema(req.structObj)
		model.ResponseSchema = schema
	}

	if req.responseType != "" {
		model.ResponseMIMEType = req.responseType
	}

	model.TopK = &req.topK
	model.TopP = &req.topP
	model.Temperature = &req.temperature

	var inputs []genai.Part

	if req.prompt == "" {
		return "", errors.New("gemini request requires input prompt")
	}
	inputs = append(inputs, genai.Text(req.prompt))

	// 이미지 존재할 경우 이미지 입력 추가
	if len(req.imageBytes) > 0 {
		inputs = append(inputs, genai.ImageData(req.imageFormat, req.imageBytes))
	}

	resp, err := model.GenerateContent(req.ctx, inputs...)
	if err != nil {
		return "", errors.New("gemini generation failed")
	}

	resultString := accumulateContent(resp.Candidates[0].Content)

	// Structured Output 모드일 경우 output 객체에 unmarshall
	if req.isStructuredOutput {
		err := JsonToStruct(resultString, req.structObj)
		if err != nil {
			return "", errors.New("gemini failed to unmarshall structured output")
		}
	}

	return resultString, nil
}

func accumulateContent(content *genai.Content) string {
	var output string
	for _, part := range content.Parts {
		if text, ok := part.(genai.Text); ok {
			output += string(text)
		}
	}
	return output
}
