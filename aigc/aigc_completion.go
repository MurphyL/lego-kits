package aigc

import (
	"log"
	"net/http"
)

/** OpenAI API - Chat Completions 格式详解 - https://zhuanlan.zhihu.com/p/692336625 */

type ChatCompletion struct {
	Messages   []CompletionMessage `json:"messages"`
	Model      string              `json:"model"`
	Stream     bool                `json:"stream"`                // 是否流式的返回消息
	ToolChoice ToolChoiceMode      `json:"tool_choice,omitempty"` // 不存在任何函数时none是默认值，否则auto是默认值
	Tools      []map[string]any    `json:"tools,omitempty"`
}

type CompletionResponse struct {
	*AgentOptions
	Stream       bool
	httpResponse *http.Response
}

type CompletionUsage struct {
	CompletionTokens uint64 // 模型生成的新token数
	PromptTokens     uint64 // 用户输入的prompt的token数
	TotalTokens      uint64 // 对话的总token数，prompt_tokens + completionTokens
}

// Ok HTTP 请求是否执行成功
func (r CompletionResponse) Ok() bool {
	return r.httpResponse.StatusCode == 200
}

// Usage 完成请求的使用统计信息
func (r CompletionResponse) Usage() *CompletionUsage {
	return &CompletionUsage{}
}

// ResolveTokens 流式返回时，token会作为的SSE(server-sent events)事件返回。http的chunk流由 data: [DONE]标记消息终止。
func (r CompletionResponse) ResolveTokens() {
	if nil != r.resolveTokens {
		r.resolveTokens("")
	} else {
		log.Println(r.httpResponse)
	}
}

func (r CompletionResponse) ResolveMessages() {
	if nil != r.resolveTokens {
		r.resolveMessages("")
	} else {
		log.Println(r.httpResponse)
	}
}

func (r CompletionResponse) ResolveTools() {
	if nil != r.resolveTokens {
		r.resolveTools("")
	} else {
		log.Println(r.httpResponse)
	}
}
