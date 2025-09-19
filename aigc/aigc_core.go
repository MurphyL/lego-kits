package aigc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/** OpenAI 的通用接口 - https://www.openaicto.com/capabilities/text-generation*/

func NewAgent(opts ...AgentOption) *Agent {
	vars := &AgentOptions{}
	for _, opt := range opts {
		opt(vars)
	}
	return &Agent{opts: vars}
}

type Agent struct {
	opts *AgentOptions
}

type AgentOptions struct {
	ServiceUrl, Token string
	CompletionPreHook func(r *http.Request)
}

type AgentOption func(*AgentOptions)

/** OpenAI API - Chat Completions 格式详解 - https://zhuanlan.zhihu.com/p/692336625 */

func (c Agent) ApplyChatCompletion(model string, messages []CompletionMessage) (*CompletionResponse, error) {
	var err error
	var request = CompletionRequest{Model: model, Messages: messages, ToolChoice: None}
	var response = new(CompletionResponse)
	// HTTP 请求
	httpRequest, err := request.MakeHttpMessage("POST", c.opts.ServiceUrl)
	if err == nil {
		httpRequest.Header.Add("Content-Type", "application/json")
		httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.opts.Token))
		if nil != c.opts.CompletionPreHook {
			c.opts.CompletionPreHook(httpRequest)
		}
		httpResponse, err := http.DefaultClient.Do(httpRequest)
		if err == nil {
			data, err := io.ReadAll(httpResponse.Body)
			if err == nil {
				json.Unmarshal(data, response)
			}
		}
	}
	return response, err
}
