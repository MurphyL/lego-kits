package aigc

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ToolChoiceMethod string

const (
	Auto     ToolChoiceMethod = "auto"
	None     ToolChoiceMethod = "none"
	Required ToolChoiceMethod = "required"
)

type MessageKind string

const (
	User   MessageKind = "user"
	System MessageKind = "system"
)

type CompletionRequest struct {
	Messages   []CompletionMessage `json:"messages"`
	Model      string              `json:"model"`
	Streaming  bool                `json:"stream"`
	ToolChoice ToolChoiceMethod    `json:"tool_choice,omitempty"`
	Tools      []map[string]any    `json:"tools,omitempty"`
}

type CompletionMessage struct {
	Role    MessageKind `json:"role"`
	Content string      `json:"content"`
}

func (request CompletionRequest) Validate() {

}

func (request CompletionRequest) MakeHttpMessage(method, url string) (*http.Request, error) {
	body, _ := json.Marshal(request)
	log.Println(string(body))
	return http.NewRequest(method, url, bytes.NewReader(body))
}

func (request CompletionRequest) AppendMessage() {

}
