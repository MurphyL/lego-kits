package aigc

// OpenAiClient - OpenAI 客户端 - 接口定义
type OpenAiClient interface {
	ApplyCompletion(text string, opts *CompletionOptions) CompletionResponse
}

type CompletionResponse interface {
	IsValid() bool
	GetResult() string
}

type CompletionOptions struct {
	Streaming bool // 是否 stream 相应
	WebSearch bool // 是否开启网络搜索
	Cleanly   bool // 是否返回说明文章
}

func ApplyCompletion(text string, opts *CompletionOptions) CompletionResponse {
	getRandomAigcAgent()
	return nil
}
