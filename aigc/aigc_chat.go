package aigc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ToolChoiceMode string

const (
	ToolChoiceModeAuto     ToolChoiceMode = "auto"
	ToolChoiceModeNone     ToolChoiceMode = "none"
	ToolChoiceModeRequired ToolChoiceMode = "required"
)

type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"      // Request
	MessageRoleSystem    MessageRole = "system"    // Request
	MessageRoleAssistant MessageRole = "assistant" // Response
)

type CompletionMessage struct {
	Role    MessageRole `json:"role"`
	Content string      `json:"content"`
}

type CompletionRequest struct {
	Messages   []CompletionMessage `json:"messages"`
	Model      string              `json:"model"`
	Streaming  bool                `json:"stream"`
	ToolChoice ToolChoiceMode      `json:"tool_choice,omitempty"`
	Tools      []map[string]any    `json:"tools,omitempty"`
}

func (request CompletionRequest) MakeHttpMessage(method, url string) (*http.Request, error) {
	body, _ := json.Marshal(request)
	return http.NewRequest(method, url, bytes.NewReader(body))
}

type CompletionResponse struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Sid     string `json:"sid"`
	Status  string `json:"status"`
	Choices []struct {
		Index   uint16 `json:"index"`
		Message struct {
			CompletionMessage
			Reason string `json:"reasoning_content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		CompletionTokens uint16 `json:"completion_tokens"`
		PromptTokens     uint16 `json:"prompt_tokens"`
		TotalTokens      uint16 `json:"total_tokens"`
	} `json:"usage"`
}
