package aigc

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

/** OpenAI 的通用接口 */

type CompletionAgent struct {
	opts *AgentOptions
}

func (c CompletionAgent) ApplyApplyCompletion(message map[string]any) (*http.Response, error) {
	body, _ := json.Marshal(message)
	log.Println(string(body))
	request, err := http.NewRequest("POST", c.opts.serviceUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(request)
	return resp, err
}

type AgentOptions struct {
	serviceUrl, token string
}

type AgentOption func(*AgentOptions)

func WithServiceProvider(url, token string) AgentOption {
	return func(opts *AgentOptions) {
		opts.serviceUrl = url
		opts.token = token
	}
}

func NewCompletionAgent(opts ...AgentOption) *CompletionAgent {
	vars := &AgentOptions{}
	for _, opt := range opts {
		opt(vars)
	}
	return &CompletionAgent{opts: vars}
}

// OpenAiClient - OpenAI 客户端 - 接口定义
type OpenAiClient interface {
	ApplyCompletion(text string, opts *AgentOptions) CompletionResponse
}

type CompletionResponse interface {
	IsValid() bool
	GetResult() string
}

func ApplyCompletion(text string, opts *AgentOptions) CompletionResponse {
	return nil
}
