package aigc

import (
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

func (c Agent) ApplyChatCompletion(model string, messages []CompletionMessage) (*http.Response, error) {
	var err error
	var httpRequest *http.Request
	var response *http.Response
	request := CompletionRequest{Model: model, Messages: messages, ToolChoice: None}
	httpRequest, err = request.MakeHttpMessage("POST", c.opts.ServiceUrl)
	if err == nil {
		if nil != c.opts.CompletionPreHook {
			c.opts.CompletionPreHook(httpRequest)
		}
		response, err = http.DefaultClient.Do(httpRequest)
	}
	return response, err
}
