package aigc

func NewOllamaAgent(url string, withOptions ...func(*AgentOptions)) Agent {
	options := &AgentOptions{
		resolveUsage: func(s string) *CompletionUsage {
			return nil
		},
	}
	for _, withOption := range withOptions {
		withOption(options)
	}
	return &ollamaAgent{internalAgent: internalAgent{Url: url, AgentOptions: options}}
}

type ollamaAgent struct {
	internalAgent
}
