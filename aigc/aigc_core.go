package aigc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/** OpenAI 的通用接口 - https://www.openaicto.com/capabilities/text-generation*/

func NewAgent(url, token string, opts ...AgentOption) *Agent {
	vars := new(AgentOptions)
	for _, opt := range opts {
		opt(vars)
	}
	return &Agent{ServiceUrl: url, Token: token}
}

type Agent struct {
	ServiceUrl, Token string
	options           *AgentOptions
}

type AgentOptions struct {
	CompletionPreHook func(r *http.Request)
}

type AgentOption func(*AgentOptions)

/** OpenAI API - Chat Completions 格式详解 - https://zhuanlan.zhihu.com/p/692336625 */

func (c Agent) ApplyChatCompletion(model string, opts ...CompletionRequestOption) (string, error) {
	var err error
	var response string
	var request = CompletionRequest{Model: model}
	for _, opt := range opts {
		opt(&request)
	}
	var body, _ = json.Marshal(request)
	httpRequest, err := http.NewRequest("POST", c.ServiceUrl, bytes.NewReader(body))
	if err == nil {
		httpRequest.Header.Add("Content-Type", "application/json")
		httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
		if nil != c.options && nil != c.options.CompletionPreHook {
			c.options.CompletionPreHook(httpRequest)
		}
		httpResponse, err := http.DefaultClient.Do(httpRequest)
		if err == nil {
			data, err := io.ReadAll(httpResponse.Body)
			if err == nil {
				response = string(data)
			}
		}
	}
	return response, err
}
