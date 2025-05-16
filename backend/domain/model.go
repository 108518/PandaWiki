package domain

import (
	"time"
)

type ModelProvider string

const (
	ModelProviderBrandOpenAI      ModelProvider = "OpenAI"
	ModelProviderBrandOllama      ModelProvider = "Ollama"
	ModelProviderBrandDeepSeek    ModelProvider = "DeepSeek"
	ModelProviderBrandMoonshot    ModelProvider = "Moonshot"
	ModelProviderBrandSiliconFlow ModelProvider = "SiliconFlow"
	ModelProviderBrandAzureOpenAI ModelProvider = "AzureOpenAI"
	ModelProviderBrandOther       ModelProvider = "Other"
)

type TokenType string

const (
	TokenTypePrompt     TokenType = "prompt"
	TokenTypeCompletion TokenType = "completion"
)

type Model struct {
	ID         string        `json:"id"`
	Provider   ModelProvider `json:"provider"`
	Model      string        `json:"model"`
	APIKey     string        `json:"api_key"`
	APIHeader  string        `json:"api_header"`
	BaseURL    string        `json:"base_url"`
	APIVersion string        `json:"api_version"` // for azure openai

	IsActive bool `json:"is_active" gorm:"default:false"`

	PromptTokens     uint64 `json:"prompt_tokens" gorm:"default:0"`
	CompletionTokens uint64 `json:"completion_tokens" gorm:"default:0"`
	TotalTokens      uint64 `json:"total_tokens" gorm:"default:0"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ModelListItem struct {
	ID         string        `json:"id"`
	Provider   ModelProvider `json:"provider"`
	Model      string        `json:"model"`
	APIKey     string        `json:"api_key"`
	APIHeader  string        `json:"api_header"`
	BaseURL    string        `json:"base_url"`
	APIVersion string        `json:"api_version"` // for azure openai

	IsActive bool `json:"is_active"`

	PromptTokens     uint64 `json:"prompt_tokens"`
	CompletionTokens uint64 `json:"completion_tokens"`
	TotalTokens      uint64 `json:"total_tokens"`
}

type ModelDetailResp struct {
	ModelListItem
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateModelReq struct {
	BaseModelInfo
}

type UpdateModelReq struct {
	ID         string `json:"id" validate:"required"`
	Model      string `json:"model" validate:"required"`
	APIKey     string `json:"api_key"`
	APIHeader  string `json:"api_header"`
	BaseURL    string `json:"base_url" validate:"required"`
	APIVersion string `json:"api_version"` // for azure openai
}

type CheckModelReq struct {
	BaseModelInfo
}

type BaseModelInfo struct {
	Provider   ModelProvider `json:"provider" validate:"required,oneof=OpenAI Ollama DeepSeek SiliconFlow Moonshot Other AzureOpenAI"`
	Model      string        `json:"model" validate:"required"`
	BaseURL    string        `json:"base_url" validate:"required"`
	APIKey     string        `json:"api_key"`
	APIHeader  string        `json:"api_header"`
	APIVersion string        `json:"api_version"` // for azure openai
}

type CheckModelResp struct {
	Error   string `json:"error"`
	Content string `json:"content"`
}

var ModelProviderBrandModelsList = map[ModelProvider][]ProviderModelListItem{
	ModelProviderBrandOpenAI: {
		{Model: "gpt-4o"},
	},
	ModelProviderBrandDeepSeek: {
		{Model: "deepseek-reasoner"},
		{Model: "deepseek-chat"},
	},
	ModelProviderBrandMoonshot: {
		{Model: "moonshot-v1-auto"},
		{Model: "moonshot-v1-8k"},
		{Model: "moonshot-v1-32k"},
		{Model: "moonshot-v1-128k"},
	},
	ModelProviderBrandAzureOpenAI: {
		{Model: "gpt-4"},
		{Model: "gpt-4o"},
		{Model: "gpt-4o-mini"},
		{Model: "gpt-4o-nano"},
		{Model: "gpt-4.1"},
		{Model: "gpt-4.1-mini"},
		{Model: "gpt-4.1-nano"},
		{Model: "o1"},
		{Model: "o1-mini"},
		{Model: "o3"},
		{Model: "o3-mini"},
		{Model: "o4-mini"},
	},
}

type GetProviderModelListReq struct {
	Provider  string `json:"provider" query:"provider" validate:"required,oneof=SiliconFlow OpenAI Ollama DeepSeek Moonshot AzureOpenAI"`
	BaseURL   string `json:"base_url" query:"base_url" validate:"required"`
	APIKey    string `json:"api_key" query:"api_key"`
	APIHeader string `json:"api_header" query:"api_header"`
}

type GetProviderModelListResp struct {
	Models []ProviderModelListItem `json:"models"`
}

type ProviderModelListItem struct {
	Model string `json:"model"`
}

type ActivateModelReq struct {
	ModelID string `json:"model_id" validate:"required"`
}
