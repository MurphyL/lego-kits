package aigc

import (
	"github.com/buger/jsonparser"
)

func NewOllamaAgent(url string, withOptions ...func(resolvers *AgentOptions)) OllamaAgent {
	options := &AgentOptions{
		resolveUsage: resolveOllamaAgentUsage,
	}
	for _, withOption := range withOptions {
		withOption(options)
	}
	agent := OllamaAgent{
		internalAgent{Url: url, AgentOptions: options},
	}
	return agent
}

type OllamaAgent struct {
	internalAgent
}

func resolveOllamaAgentUsage(d []byte) *CompletionUsage {
	ctn, _ := jsonparser.GetInt(d, "eval_count")
	ptn, _ := jsonparser.GetInt(d, "prompt_eval_count")
	return &CompletionUsage{
		CompletionTokens: uint64(ctn),
		PromptTokens:     uint64(ptn),
		TotalTokens:      uint64(ctn + ptn),
	}
}
