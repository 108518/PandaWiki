package domain

import (
	"time"

	modelkitConsts "github.com/chaitin/ModelKit/consts"
	modelkitDomain "github.com/chaitin/ModelKit/domain"
)

type ModelProvider string

const (
	ModelProviderBrandOpenAI      ModelProvider = "OpenAI"
	ModelProviderBrandOllama      ModelProvider = "Ollama"
	ModelProviderBrandDeepSeek    ModelProvider = "DeepSeek"
	ModelProviderBrandMoonshot    ModelProvider = "Moonshot"
	ModelProviderBrandSiliconFlow ModelProvider = "SiliconFlow"
	ModelProviderBrandAzureOpenAI ModelProvider = "AzureOpenAI"
	ModelProviderBrandBaiZhiCloud ModelProvider = "BaiZhiCloud"
	ModelProviderBrandHunyuan     ModelProvider = "Hunyuan"
	ModelProviderBrandBaiLian     ModelProvider = "BaiLian"
	ModelProviderBrandVolcengine  ModelProvider = "Volcengine"
	ModelProviderBrandGemini      ModelProvider = "Gemini"
	ModelProviderBrandZhiPu       ModelProvider = "ZhiPu" // 智谱
	ModelProviderBrandOther       ModelProvider = "Other"
)

type ModelType string

const (
	ModelTypeChat      ModelType = "chat"
	ModelTypeEmbedding ModelType = "embedding"
	ModelTypeRerank    ModelType = "rerank"
)

type Model struct {
	ID         string        `json:"id"`
	Provider   ModelProvider `json:"provider"`
	Model      string        `json:"model"`
	APIKey     string        `json:"api_key"`
	APIHeader  string        `json:"api_header"`
	BaseURL    string        `json:"base_url"`
	APIVersion string        `json:"api_version"` // for azure openai
	Type       ModelType     `json:"type" gorm:"default:chat;uniqueIndex"`

	IsActive bool `json:"is_active" gorm:"default:false"`

	PromptTokens     uint64 `json:"prompt_tokens" gorm:"default:0"`
	CompletionTokens uint64 `json:"completion_tokens" gorm:"default:0"`
	TotalTokens      uint64 `json:"total_tokens" gorm:"default:0"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToModelkitModel converts domain.Model to modelkitDomain.PandaModel
func (m *Model) ToModelkitModel() (*modelkitDomain.ModelMetadata, error) {
	provider := modelkitConsts.ParseModelProvider(string(m.Provider))
	modelType, err := modelkitConsts.ParseModelType(string(m.Type))
	if err != nil {
		return nil, err
	}
	return &modelkitDomain.ModelMetadata{
		Provider:   provider,
		ModelName:  m.Model,
		APIKey:     m.APIKey,
		BaseURL:    m.BaseURL,
		APIVersion: m.APIVersion,
		APIHeader:  m.APIHeader,
		ModelType:  modelType,
	}, nil
}

type ModelListItem struct {
	ID         string        `json:"id"`
	Provider   ModelProvider `json:"provider"`
	Model      string        `json:"model"`
	APIKey     string        `json:"api_key"`
	APIHeader  string        `json:"api_header"`
	BaseURL    string        `json:"base_url"`
	APIVersion string        `json:"api_version"` // for azure openai
	Type       ModelType     `json:"type"`

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
	ID string `json:"id" validate:"required"`
	BaseModelInfo
}

type CheckModelReq struct {
	BaseModelInfo
}

type BaseModelInfo struct {
	Provider   ModelProvider `json:"provider" validate:"required,oneof=OpenAI Ollama DeepSeek SiliconFlow Moonshot Other AzureOpenAI BaiZhiCloud Hunyuan BaiLian Volcengine Gemini ZhiPu"`
	Model      string        `json:"model" validate:"required"`
	BaseURL    string        `json:"base_url" validate:"required"`
	APIKey     string        `json:"api_key"`
	APIHeader  string        `json:"api_header"`
	APIVersion string        `json:"api_version"` // for azure openai
	Type       ModelType     `json:"type" validate:"required,oneof=chat embedding rerank"`
}

type CheckModelResp struct {
	Error   string `json:"error"`
	Content string `json:"content"`
}

type GetProviderModelListReq struct {
	Provider  string    `json:"provider" query:"provider" validate:"required,oneof=SiliconFlow OpenAI Ollama DeepSeek Moonshot AzureOpenAI BaiZhiCloud Hunyuan BaiLian Volcengine Gemini ZhiPu"`
	BaseURL   string    `json:"base_url" query:"base_url" validate:"required"`
	APIKey    string    `json:"api_key" query:"api_key"`
	APIHeader string    `json:"api_header" query:"api_header"`
	Type      ModelType `json:"type" query:"type" validate:"required,oneof=chat embedding rerank"`
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
