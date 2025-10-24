package aigc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Agent 客户端 OpenAI 的通用接口 - https:www.openaicto.com/capabilities/text-generation
type Agent interface {
	ApplyCompletion(model string, withOptions ...func(*ChatCompletion)) (*CompletionResponse, error)
}

type AgentOptions struct {
	resolveTokens   func(string)
	resolveMessages func(string)
	resolveTools    func(string)
	resolveUsage    func(string) *CompletionUsage
}

type internalAgent struct {
	*AgentOptions
	Url   string
	Token string
}

func (c internalAgent) ApplyCompletion(model string, withOptions ...func(*ChatCompletion)) (*CompletionResponse, error) {
	var chatCompletion = ChatCompletion{Model: model}
	for _, withOption := range withOptions {
		withOption(&chatCompletion)
	}
	var body, _ = json.Marshal(chatCompletion)
	if httpRequest, err := http.NewRequest("POST", c.Url, bytes.NewReader(body)); err == nil {
		httpRequest.Header.Add("Content-Type", "application/json")
		httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
		if httpResponse, re := http.DefaultClient.Do(httpRequest); re == nil {
			return &CompletionResponse{
				Stream:       chatCompletion.Stream,
				httpResponse: httpResponse,
				AgentOptions: c.AgentOptions,
			}, nil
		} else {
			return nil, fmt.Errorf("执行HTTP请求出错：%s", re.Error())
		}
	} else {
		return nil, fmt.Errorf("构造HTTP请求出错：%s", err.Error())
	}
}
