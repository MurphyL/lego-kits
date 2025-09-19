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
